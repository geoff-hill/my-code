package au.com.patientzero.demos.edge;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.http.HttpHeaders;
import org.springframework.messaging.rsocket.RSocketRequester;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

@SpringBootApplication
public class EdgeDemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(EdgeDemoApplication.class, args);
	}

	@Bean
	RouteLocator gateway(RouteLocatorBuilder rlb) {
		return rlb
				.routes()
				.route( rs -> rs
						.path("/proxy")
						.filters( fs -> fs
								.setPath("/customers")
								.setResponseHeader(HttpHeaders.ACCESS_CONTROL_ALLOW_ORIGIN, "*")
						)
						.uri("http://localhost:8900")
				)
				.build();
	}

	@Bean
	WebClient webClient(WebClient.Builder builder) {
		return builder.build();
	}

	@Bean
	RSocketRequester rSocketRequester(RSocketRequester.Builder builder) {
		return builder.tcp("localhost", 9900);
	}
}

@Component
class CrmClient {
	private final WebClient        http;
	private final RSocketRequester rsocket;

	CrmClient(WebClient http, RSocketRequester rsocket) {
		this.http = http;
		this.rsocket = rsocket;
	}

	Flux<Customer> getCustomers() {
		return this.http
				.get()
				.uri("http://localhost:8900/customers")
				.retrieve()
				.bodyToFlux(Customer.class);

	}

	Mono<Profile> getProfileFor(Integer customerId) {

		return this.rsocket.route("profiles.{cid}", customerId).retrieveMono(Profile.class);
	}

	Flux<CustomerProfile> getCustomerProfiles(){
		return this.getCustomers()
				.flatMap( customer -> Mono.zip(Mono.just(customer), this.getProfileFor(customer.id()))
				)
				.map( tuple2 -> new CustomerProfile(tuple2.getT1(), tuple2.getT2()));
	}
}

@Controller
@ResponseBody
class CrmRestController {
	private final CrmClient crm;

	CrmRestController(CrmClient crm) {
		this.crm = crm;
	}

	@GetMapping ("/cps")
	Flux<CustomerProfile> get() {
		return this.crm.getCustomerProfiles();
	}
}

//@Controller
//class CrmGraphqlController {
//	private final CrmClient crm;
//
//	CrmGraphqlController(CrmClient crm) {
//		this.crm = crm;
//	}
//
//	@QueryMapping
//	Flux<Customer> customers() {
//		return this.crm.getCustomers();
//	}
//
//	@SchemaMapping (typeName="Customer")
//	Mono<Profile> profile (Customer customer) {
//		return this.crm.getProfileFor(customer.id());
//	}
//}

record CustomerProfile(Customer customer, Profile profile){}
record Profile (Integer id, Integer customerId){}
record Customer (Integer id, String name){}
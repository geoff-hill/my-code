package au.com.patientzero.demos.customers;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.actuate.info.Info;
import org.springframework.boot.actuate.info.InfoContributor;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.condition.ConditionalOnCloudPlatform;
import org.springframework.boot.cloud.CloudPlatform;
import org.springframework.context.annotation.Bean;
import org.springframework.data.annotation.Id;
import org.springframework.data.repository.reactive.ReactiveCrudRepository;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.ResponseBody;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.Collections;




@SpringBootApplication
public class CustomersDemoApplication {
	@Bean
	@ConditionalOnCloudPlatform(CloudPlatform.KUBERNETES)
	ApplicationRunner runner() {
		return args -> System.out.println("Hello, kubernetes world");
	}

	public static void main(String[] args) {
		Logger logger = LoggerFactory.getLogger(CustomersDemoApplication.class);
		logger.info(String.valueOf(Runtime.getRuntime().availableProcessors()));
		SpringApplication.run(CustomersDemoApplication.class, args);
	}

}

@Component
class MyInfoContributor implements InfoContributor {

	@Override
	public void contribute(Info.Builder builder) {
		String processors = String.valueOf(Runtime.getRuntime().availableProcessors());
		builder.withDetail("ergonomics", Collections.singletonMap("availableProcessors", processors));
	}

}

@Controller
@ResponseBody
class CustomerRestController {
	private final CustomerRepository repository;

	CustomerRestController(CustomerRepository repository) {
		this.repository = repository;
	}

	@GetMapping("/customers")
	Flux<Customer> get() {
		return repository.findAll();
	}

	@GetMapping("/customers/{id}")
	Mono<Customer> get(@PathVariable Integer id) { return repository.findById(id); }
}
interface CustomerRepository extends ReactiveCrudRepository<Customer, Integer> {}
record Customer (@Id Integer id, String name){}


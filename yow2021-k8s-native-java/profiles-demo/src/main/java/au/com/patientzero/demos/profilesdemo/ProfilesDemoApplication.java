package au.com.patientzero.demos.profilesdemo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.messaging.handler.annotation.DestinationVariable;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.stereotype.Controller;

import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.IntStream;

@SpringBootApplication
public class ProfilesDemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(ProfilesDemoApplication.class, args);
	}

}

@Controller
class ProfileRSocketController {
	ConcurrentHashMap<Integer, Profile> db = new ConcurrentHashMap<>();

	ProfileRSocketController() {
		IntStream.rangeClosed(1, 7).boxed()
			.map((Integer x) -> new Profile(x,x))
			.toList()
			.forEach((p) -> db.put(p.customerId(), p));
	}

	@MessageMapping("profiles.{customerId}")
	Profile getProfileFor(@DestinationVariable Integer customerId) {
		return db.get(customerId);
	}
}

record Profile(Integer id, Integer customerId) {}

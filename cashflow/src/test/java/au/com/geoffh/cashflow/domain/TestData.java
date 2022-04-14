package au.com.geoffh.cashflow.domain;

import com.github.javafaker.Faker;

import java.math.BigDecimal;
import java.util.Date;
import java.util.Random;
import java.util.UUID;

public class TestData extends Faker {
    public static TestData testData() {
        Random random = new Random(new Date().getTime());
        return new TestData(random);
    }

    public TestData(Random random) {
        super(random);
    }

    public UUID uuid() {
        return UUID.randomUUID();
    }

    public BigDecimal dollars() {
        return dollars(4);
    }

    public BigDecimal dollars(int maxDollarDigits) {
        String template = "#".repeat(maxDollarDigits) + ".##";
        String numerified = numerify(template);
        return new BigDecimal(numerified);
    }
}

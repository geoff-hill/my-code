package au.com.geoffh.cashflow.domain;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.util.Objects;
import java.util.UUID;

public class Bill {
    private final UUID id;
    private final String name;
    private final BigDecimal amount;
    private final BillingFrequency billingFrequency;
    private final LocalDate nextDue;


    public Bill(UUID id, String name, BigDecimal amount, BillingFrequency billingPeriod, LocalDate nextDue) {
        this.id = id;
        this.name = name;
        this.amount = amount;
        this.billingFrequency = billingPeriod;
        this.nextDue = nextDue;
    }

    public Bill(String name, BigDecimal amount, BillingFrequency billingPeriod, LocalDate nextDue) {
        this(UUID.randomUUID(), name, amount, billingPeriod, nextDue);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Bill bill = (Bill) o;
        return id.equals(bill.id);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id);
    }

    public UUID getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public BigDecimal getAmount() {
        return amount;
    }

    public BillingFrequency getBillingFrequency() {
        return billingFrequency;
    }

    public LocalDate getNextDue() {
        return nextDue;
    }
}

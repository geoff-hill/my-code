package au.com.geoffh.cashflow.domain;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.util.Objects;

public class PayBucket {
    private final LocalDate payday;
    private final BigDecimal totalIncome;

    public PayBucket(LocalDate payday) {
        this.payday = payday;
        totalIncome = BigDecimal.ZERO;
    }

    public PayBucket(LocalDate payday, BigDecimal totalIncome) {
        this.payday = payday;
        this.totalIncome = totalIncome;
    }

    public BigDecimal getTotalIncome() {
        return totalIncome;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        PayBucket payBucket = (PayBucket) o;
        return Objects.equals(payday, payBucket.payday);
    }

    @Override
    public int hashCode() {
        return Objects.hash(payday);
    }

    public LocalDate getPayday() {
        return payday;
    }
}

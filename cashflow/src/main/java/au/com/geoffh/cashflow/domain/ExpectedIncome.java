package au.com.geoffh.cashflow.domain;

import java.math.BigDecimal;
import java.time.LocalDate;

public class ExpectedIncome {
    private final LocalDate expectedPayday;
    private final BigDecimal amount;

    public ExpectedIncome(LocalDate expectedPayday, BigDecimal amount) {
        this.expectedPayday = expectedPayday;
        this.amount = amount;
    }

    public LocalDate getExpectedPayday() {
        return expectedPayday;
    }

    public BigDecimal getAmount() {
        return amount;
    }
}

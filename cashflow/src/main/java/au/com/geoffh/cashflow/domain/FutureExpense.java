package au.com.geoffh.cashflow.domain;

import java.math.BigDecimal;
import java.time.LocalDate;

public class FutureExpense {
    private final LocalDate dueDate;
    private final LocalDate targetPayDate;
    private final BigDecimal amount;

    public FutureExpense(LocalDate dueDate, BigDecimal amount) {
        this.dueDate = dueDate;
        this.amount = amount;
        targetPayDate = dueDate;
    }

    public LocalDate getDueDate() {
        return dueDate;
    }

    public LocalDate getTargetPayDate() {
        return targetPayDate;
    }

    public BigDecimal getAmount() {
        return amount;
    }
}

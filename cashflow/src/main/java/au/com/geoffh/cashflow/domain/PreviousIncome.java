package au.com.geoffh.cashflow.domain;

import java.time.LocalDate;

public class PreviousIncome {
    private final LocalDate receivedDate;

    public PreviousIncome(LocalDate receivedDate) {
        this.receivedDate = receivedDate;
    }

    public LocalDate getReceivedDate() {
        return receivedDate;
    }
}

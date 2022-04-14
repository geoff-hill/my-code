package au.com.geoffh.cashflow.ports;

import au.com.geoffh.cashflow.domain.ExpectedIncome;

import java.time.LocalDate;
import java.util.List;

public interface ExpectedIncomeRepository {
    public List<ExpectedIncome> getExpectedIncome(LocalDate asAtDate, LocalDate horizonDate);
}

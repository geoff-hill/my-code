package au.com.geoffh.cashflow.ports;

import au.com.geoffh.cashflow.domain.ExpectedIncome;
import au.com.geoffh.cashflow.domain.FutureExpense;

import java.time.LocalDate;
import java.util.List;

public interface FutureExpenseRepository {
    public List<FutureExpense> getFutureExpenses(LocalDate asAtDate, LocalDate horizonDate);
}

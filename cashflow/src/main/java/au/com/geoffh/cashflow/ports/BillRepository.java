package au.com.geoffh.cashflow.ports;

import au.com.geoffh.cashflow.domain.Bill;

import java.util.List;

public interface BillRepository {
    List<Bill> getActiveBills();
}

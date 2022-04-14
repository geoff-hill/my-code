package au.com.geoffh.cashflow.domain;

import au.com.geoffh.cashflow.ports.BillRepository;

import java.util.List;

public class BillSchedule {
    private final BillRepository billRepository;

    public BillSchedule(BillRepository billRepository) {
        this.billRepository = billRepository;
    }

    public List<Bill> getActiveBills() {
        return billRepository.getActiveBills();
    }
}

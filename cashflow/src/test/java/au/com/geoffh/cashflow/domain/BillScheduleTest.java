package au.com.geoffh.cashflow.domain;

import au.com.geoffh.cashflow.ports.BillRepository;
import org.junit.jupiter.api.Test;

import java.time.LocalDate;
import java.util.ArrayList;
import java.util.List;

import static java.time.LocalDate.now;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class BillScheduleTest {
    private static final TestData testData = TestData.testData();
    private static final LocalDate TODAY = now();

    private static BillingFrequency randomBillingFrequency() {
        int n = BillingFrequency.values().length;
        int random = testData.number().numberBetween(0, n);
        return BillingFrequency.values()[random];
    }

    private static Bill newTestBill() {
        return new Bill(
                testData.uuid(),
                testData.company().name(),
                testData.dollars(),
                randomBillingFrequency(),
                TODAY.plusDays(testData.number().numberBetween(1, 365))
        );
    }

    private List<Bill> newTestBills(int num) {
        var ret = new ArrayList<Bill>(num);
        for (int i = 0; i < num; i++) {
            ret.add(newTestBill());
        }
        return ret;
    }

    private static void assertAllFieldsEqual(Bill expected, Bill returned) {
        assertEquals(expected.getId(), returned.getId());
        assertEquals(expected.getAmount(), returned.getAmount());
        assertEquals(expected.getBillingFrequency(), returned.getBillingFrequency());
        assertEquals(expected.getName(), returned.getName());
        assertEquals(expected.getNextDue(), returned.getNextDue());
    }

    private static void assertListsOfBillsSame(List<Bill> expected, List<Bill> ret) {
        assertEquals(expected.size(), ret.size());
        for (int i = 0; i < ret.size(); i++) {
            expected.get(i);
            var returned = ret.get(i);
        }
    }

    @Test
    void getActiveBills_shouldCallBillRepoAndPassBackAllBillsReturned() {
        // given
        var billRepo = mock(BillRepository.class);
        var testBills = newTestBills(5);
        when(billRepo.getActiveBills()).thenReturn(testBills);
        var billSchedule = new BillSchedule(billRepo);
        // when
        var ret = billSchedule.getActiveBills();
        // then
        assertListsOfBillsSame(testBills, ret);
    }

}
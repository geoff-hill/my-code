package au.com.geoffh.cashflow.domain;

import au.com.geoffh.cashflow.ports.BillRepository;
import au.com.geoffh.cashflow.ports.ExpectedIncomeRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.util.List;
import java.util.stream.Stream;

import static java.time.LocalDate.now;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class PayScheduleTest {

    // declare some handy dates
    private static final LocalDate TODAY = now(); // make sure we always have the same now()

    @Test
    void getPayBuckets_shouldReturnPayBucketsForTheDatesReturnedFromTheExpectedIncomeRepo() {
        // given
        var testIncome = Stream.of(TODAY.plusDays(1), TODAY.plusDays(5)).map(dt -> new ExpectedIncome(dt, BigDecimal.ZERO)).toList();
        var incomeRepo = mock(ExpectedIncomeRepository.class);
        when(incomeRepo.getExpectedIncome(TODAY, TODAY.plusDays(7))).thenReturn(testIncome);
        PaySchedule schedule = new PaySchedule(incomeRepo);
        // when
        var ret = schedule.getExpectedPayBuckets(TODAY, TODAY.plusDays(7));
        // then
        assertEquals(testIncome.size(), ret.size());
        assertEquals(testIncome.get(0).getExpectedPayday(), ret.get(0).getPayday());
        assertEquals(testIncome.get(1).getExpectedPayday(), ret.get(1).getPayday());
    }

    @Test
    void getPayBuckets_shouldMergeMultiplePaysOnTheSameDay() {
        // given
        var incomeRepo = mock(ExpectedIncomeRepository.class);
        var testIncome = Stream.of(TODAY.plusDays(1), TODAY.plusDays(3), TODAY.plusDays(3), TODAY.plusDays(5)).map(expectedPayday -> new ExpectedIncome(expectedPayday, BigDecimal.ZERO)).toList();
        when(incomeRepo.getExpectedIncome(TODAY, TODAY.plusDays(7))).thenReturn(testIncome);
        BillRepository billRepo = mock(BillRepository.class);
        PaySchedule schedule = new PaySchedule(incomeRepo);// when
        // when
        var ret = schedule.getExpectedPayBuckets(TODAY, TODAY.plusDays(7));
        // then
        assertEquals(3, ret.size());
        assertEquals(testIncome.get(0).getExpectedPayday(), ret.get(0).getPayday());
        assertEquals(testIncome.get(1).getExpectedPayday(), ret.get(1).getPayday());
        assertEquals(testIncome.get(3).getExpectedPayday(), ret.get(2).getPayday());
    }

    @Test
    void getPayBuckets_shouldHaveTheCorrectPayAmountsForEachBucket() {
        // given
        var testIncome = List.of(
                new ExpectedIncome(TODAY.plusDays(1), new BigDecimal("1276.25")),
                new ExpectedIncome(TODAY.plusDays(3), new BigDecimal("1276.25")),
                new ExpectedIncome(TODAY.plusDays(3), new BigDecimal("1000")),
                new ExpectedIncome(TODAY.plusDays(5), new BigDecimal("1276.25"))
        );
        var incomeRepo = mock(ExpectedIncomeRepository.class);
        when(incomeRepo.getExpectedIncome(TODAY, TODAY.plusDays(7))).thenReturn(testIncome);
        PaySchedule schedule = new PaySchedule(incomeRepo);// when
        // when
        var ret = schedule.getExpectedPayBuckets(TODAY, TODAY.plusDays(7));
        // then
        assertEquals(3, ret.size());
        assertEquals(TODAY.plusDays(1), ret.get(0).getPayday());
        assertEquals(new BigDecimal("1276.25"), ret.get(0).getTotalIncome());
        assertEquals(TODAY.plusDays(3), ret.get(1).getPayday());
        assertEquals(new BigDecimal("2276.25"), ret.get(1).getTotalIncome());
        assertEquals(TODAY.plusDays(5), ret.get(2).getPayday());
        assertEquals(new BigDecimal("1276.25"), ret.get(2).getTotalIncome());
    }

}
package au.com.geoffh.cashflow.domain;

import au.com.geoffh.cashflow.ports.ExpectedIncomeRepository;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.util.List;

import static java.util.Comparator.comparing;
import static java.util.stream.Collectors.groupingBy;
import static java.util.stream.Collectors.reducing;

public class PaySchedule {
    private final ExpectedIncomeRepository futureIncomeRepository;

    public PaySchedule(ExpectedIncomeRepository futureIncomeRepository) {
        this.futureIncomeRepository = futureIncomeRepository;
    }

    public List<PayBucket> getExpectedPayBuckets(LocalDate currentDate, LocalDate lastDate) {
        return futureIncomeRepository
                .getExpectedIncome(currentDate, lastDate)
                .stream()
                .collect(groupingBy(ExpectedIncome::getExpectedPayday,
                        reducing(BigDecimal.ZERO, ExpectedIncome::getAmount, BigDecimal::add)))
                .entrySet()
                .stream()
                .map((e) -> new PayBucket(e.getKey(), e.getValue()))
                .sorted(comparing(PayBucket::getPayday))
                .toList();
    }

}

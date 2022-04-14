package au.com.patientzero.cheeseria.models;

import java.math.BigDecimal;

public record Cheese(Integer id, String title, BigDecimal price, String description, String category, String image ) {

}

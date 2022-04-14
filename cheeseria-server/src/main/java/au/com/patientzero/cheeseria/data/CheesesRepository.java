package au.com.patientzero.cheeseria.data;

import java.util.Collection;

import au.com.patientzero.cheeseria.models.Cheese;

public interface CheesesRepository {
    Collection<Cheese> getAll();  
}

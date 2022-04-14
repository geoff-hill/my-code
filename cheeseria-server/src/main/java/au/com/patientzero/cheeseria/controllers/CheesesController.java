package au.com.patientzero.cheeseria.controllers;

import au.com.patientzero.cheeseria.data.CheesesRepository;
import au.com.patientzero.cheeseria.models.Cheese;
import io.javalin.http.Context;
import io.javalin.plugin.openapi.annotations.HttpMethod;
import io.javalin.plugin.openapi.annotations.OpenApi;
import io.javalin.plugin.openapi.annotations.OpenApiContent;
import io.javalin.plugin.openapi.annotations.OpenApiResponse;

public class CheesesController {
  private final CheesesRepository cheesesRepository;

  public CheesesController(CheesesRepository cheesesRepository) {
    this.cheesesRepository = cheesesRepository;
  }
  
  @OpenApi(
    summary = "Get all cheeses",
    description = "Gets all the cheeses from the repository",
    operationId = "getAllCheeses",
    path = "/cheeses",
    method = HttpMethod.GET,
    tags = {"Cheese"},
    responses = {
      @OpenApiResponse(status = "200", content = {@OpenApiContent(from = Cheese.class, isArray = true)})
    })
  public void getAll(Context ctx) {
    ctx.json(cheesesRepository.getAll());
  }
}

package endpoint;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import model.IPDatabase;

import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class CountryHandler implements HttpHandler {
    @Override
    public void handle(HttpExchange exchange) throws IOException {
        System.out.println("got country request");

        var is = exchange.getRequestBody();
        var request = new String(is.readAllBytes(), StandardCharsets.UTF_8);
        String response;

        exchange.getResponseHeaders().add("Content-Type", "text/plain");

        try {
            response = IPDatabase.getInstance().getCountry(request);
            exchange.sendResponseHeaders(200, response.length());
        } catch (NoSuchFieldException | RuntimeException e) {
            response = "<h1>400 Bad Request</h1>" + e.getMessage();
            exchange.sendResponseHeaders(400, response.length());
        }

        var os = exchange.getResponseBody();
        os.write(response.getBytes(StandardCharsets.UTF_8));
        os.close();
    }
}

import com.sun.net.httpserver.HttpServer;
import endpoint.CoordHandler;
import endpoint.CountryHandler;
import model.IPDatabase;
import model.IPEntry;
import picocli.CommandLine;
import utils.Arguments;

import java.io.*;
import java.net.InetSocketAddress;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) {
        var arguments = new Arguments();
        var cli = new CommandLine(arguments);
        String inputFile = null;

        try {
            cli.parseArgs(args);
            inputFile = arguments.getFile();
        } catch (CommandLine.ParameterException e) {
            var writer = cli.getErr();

            writer.write(e.getMessage() + "\n");
            cli.usage(writer);
            writer.flush();

            System.exit(0);
        }

        if (inputFile == null) {
            System.out.println(cli.getUsageMessage());
            System.exit(-1);
        }

        System.out.println("Started parsing InputFile: "+inputFile);

        List<IPEntry> ipData = new ArrayList<>();

        long startTime = System.currentTimeMillis();

        try (var input = new BufferedReader(new FileReader(inputFile))) {
            String line;

            while ( (line = input.readLine()) != null ) {
                ipData.add(IPEntry.parseLine(line));
            }

        } catch (FileNotFoundException e) {
            System.out.println(inputFile+" not found!");
        } catch (ParseException | IOException e) {
            e.printStackTrace();
        }

        long parsingTime = System.currentTimeMillis() - startTime;

        IPDatabase.initialize(ipData);

        System.out.println("Finished parsing");
        System.out.println("---");
        System.out.println("number of entries: "+ipData.size());
        System.out.println("Parsing time: "+parsingTime+"ms");
        System.out.println("---");
        System.out.println("Starting HTTP server...");

        startServer(arguments.getPort());
    }

    private static void startServer(int port) {
        try {
            var httpServer = HttpServer.create(new InetSocketAddress(port), 0);

            System.out.println("Webserver started on port "+port);

            httpServer.createContext("/country", new CountryHandler());
            httpServer.createContext("/coords", new CoordHandler());
            httpServer.start();

        } catch (IOException e) {
            System.out.println("Failed to start webserver!");
            e.printStackTrace();
        }
    }
}

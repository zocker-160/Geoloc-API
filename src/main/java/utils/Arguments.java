package utils;

import picocli.CommandLine;

@CommandLine.Command(mixinStandardHelpOptions = true, helpCommand = true)
public class Arguments {
    @CommandLine.Parameters(paramLabel = "File", arity = "1", description = "IP database file")
    private String file;

    @CommandLine.Option(names = {"-p", "--port"}, description = "port of webserver")
    private int port = 9000;

    public String getFile() {
        return file;
    }

    public int getPort() {
        return port;
    }
}

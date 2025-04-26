
abstract class DataMiner {
    // Template method that defines the algorithm structure
    public final void mine(String path) {
        String data = extractData(path);
        String text = parseData(data);
        String analysis = analyze(text);
        sendReport(analysis);
    }

    // These methods will be implemented by subclasses
    protected abstract String extractData(String path);

    protected abstract String parseData(String data);

    // These methods are common to all subclasses
    protected String analyze(String text) {
        System.out.println("Analyzing data...");
        return "Analysis results for: " + text;
    }

    protected void sendReport(String analysis) {
        System.out.println("Sending report: " + analysis);
    }
}

class PDFDataMiner extends DataMiner {
    @Override
    protected String extractData(String path) {
        System.out.println("Extracting data from PDF: " + path);
        return "PDF content";
    }

    @Override
    protected String parseData(String data) {
        System.out.println("Parsing PDF data");
        return "Parsed PDF content";
    }
}

class DocDataMiner extends DataMiner {
    @Override
    protected String extractData(String path) {
        System.out.println("Extracting data from DOC: " + path);
        return "DOC content";
    }

    @Override
    protected String parseData(String data) {
        System.out.println("Parsing DOC data");
        return "Parsed DOC content";
    }
}

public class TemplateMethodDemo {
    public static void main(String[] args) {
        System.out.println("Mining PDF document:");
        DataMiner pdfMiner = new PDFDataMiner();
        pdfMiner.mine("path/to/document.pdf");

        System.out.println("\nMining DOC document:");
        DataMiner docMiner = new DocDataMiner();
        docMiner.mine("path/to/document.doc");
    }
}
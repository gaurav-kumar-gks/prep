import java.util.List;
import java.util.ArrayList;
import java.util.Collections;

interface Strategy {
    List<String> doAlgorithm(List<String> data);
}

class ConcreteStrategyA implements Strategy {
    @Override
    public List<String> doAlgorithm(List<String> data) {
        List<String> result = new ArrayList<>(data);
        Collections.sort(result);
        return result;
    }
}

class ConcreteStrategyB implements Strategy {
    @Override
    public List<String> doAlgorithm(List<String> data) {
        List<String> result = new ArrayList<>(data);
        Collections.sort(result, Collections.reverseOrder());
        return result;
    }
}

class Context {
    private Strategy strategy;

    public Context(Strategy strategy) {
        this.strategy = strategy;
    }

    public void setStrategy(Strategy strategy) {
        this.strategy = strategy;
    }

    public Strategy getStrategy() {
        return strategy;
    }

    public void doSomeBusinessLogic() {
        System.out.println("Context: Sorting data using the strategy (not sure how it'll do it)");
        List<String> data = List.of("a", "b", "c", "d", "e");
        List<String> result = strategy.doAlgorithm(data);
        System.out.println(String.join(",", result));
    }
}

public class StrategyDemo {
    public static void main(String[] args) {
        Context context = new Context(new ConcreteStrategyA());
        System.out.println("Client: Strategy is set to " + context.getStrategy().getClass().getSimpleName());
        context.doSomeBusinessLogic();
        System.out.println();

        System.out.println("Client: Strategy is set to reverse sorting.");
        context.setStrategy(new ConcreteStrategyB());
        context.doSomeBusinessLogic();
    }
}
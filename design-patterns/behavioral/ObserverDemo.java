import java.util.ArrayList;
import java.util.List;
import java.util.Random;

interface Publisher {
    void subscribe(Subscriber subscriber);

    void unsubscribe(Subscriber subscriber);

    void notifySubscribers();
}

interface Subscriber {
    void update(Publisher publisher);
}

class ConcretePublisher implements Publisher {
    private int state = 0;
    private List<Subscriber> subscribers = new ArrayList<>();
    private Random random = new Random();

    public int getState() {
        return state;
    }

    public void setState(int state) {
        this.state = state;
    }

    @Override
    public void subscribe(Subscriber subscriber) {
        System.out.println("Publisher: Added a subscriber.");
        subscribers.add(subscriber);
    }

    @Override
    public void unsubscribe(Subscriber subscriber) {
        subscribers.remove(subscriber);
    }

    @Override
    public void notifySubscribers() {
        System.out.println("Subject: Notifying subscribers...");
        for (Subscriber subscriber : subscribers) {
            subscriber.update(this);
        }
    }

    public void someBusinessLogic() {
        System.out.println("Publisher: I'm doing something important.");
        this.state = random.nextInt(10);

        System.out.println("Publisher: My state has just changed to: " + this.state);
        this.notifySubscribers();
    }
}

class ConcreteSubscriberA implements Subscriber {
    @Override
    public void update(Publisher publisher) {
        System.out.println("ConcreteSubscriberA updated");
    }
}

class ConcreteSubscriberB implements Subscriber {
    @Override
    public void update(Publisher publisher) {
        System.out.println("ConcreteSubscriberB updated");
    }
}

public class ObserverDemo {
    public static void main(String[] args) {
        ConcretePublisher publisher = new ConcretePublisher();

        Subscriber subscriberA = new ConcreteSubscriberA();
        publisher.subscribe(subscriberA);

        Subscriber subscriberB = new ConcreteSubscriberB();
        publisher.subscribe(subscriberB);

        publisher.someBusinessLogic();
        publisher.someBusinessLogic();

        publisher.unsubscribe(subscriberA);

        publisher.someBusinessLogic();
    }
}
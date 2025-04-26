import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

abstract class State {
    protected VendingMachine vendingMachine;

    public State(VendingMachine vendingMachine) {
        this.vendingMachine = vendingMachine;
    }

    public abstract void selectItem();

    public abstract void addItem(String item);

    public abstract void addMoney(double amount);

    public abstract void dispenseItem();
}

class HasItemState extends State {
    public HasItemState(VendingMachine vendingMachine) {
        super(vendingMachine);
    }

    @Override
    public void selectItem() {
        System.out.println("Item selected.");
        vendingMachine.setState(vendingMachine.getItemRequestedState());
    }

    @Override
    public void addItem(String item) {
        System.out.println("Item already present. Cannot add another item.");
    }

    @Override
    public void addMoney(double amount) {
        System.out.println("Please select an item first.");
    }

    @Override
    public void dispenseItem() {
        System.out.println("Please select an item first.");
    }
}

class ItemRequestedState extends State {
    public ItemRequestedState(VendingMachine vendingMachine) {
        super(vendingMachine);
    }

    @Override
    public void selectItem() {
        System.out.println("Item already selected.");
    }

    @Override
    public void addItem(String item) {
        System.out.println("Cannot add item. Item already selected.");
    }

    @Override
    public void addMoney(double amount) {
        System.out.println("Money added: " + amount);
        vendingMachine.setState(vendingMachine.getHasMoneyState());
    }

    @Override
    public void dispenseItem() {
        System.out.println("Please add money first.");
    }
}

class HasMoneyState extends State {
    public HasMoneyState(VendingMachine vendingMachine) {
        super(vendingMachine);
    }

    @Override
    public void selectItem() {
        System.out.println("Item already selected.");
    }

    @Override
    public void addItem(String item) {
        System.out.println("Cannot add item. Item already selected.");
    }

    @Override
    public void addMoney(double amount) {
        System.out.println("Money already added.");
    }

    @Override
    public void dispenseItem() {
        System.out.println("Dispensing item...");
        vendingMachine.setState(vendingMachine.getNoItemState());
    }
}

class NoItemState extends State {
    public NoItemState(VendingMachine vendingMachine) {
        super(vendingMachine);
    }

    @Override
    public void selectItem() {
        System.out.println("No item available.");
    }

    @Override
    public void addItem(String item) {
        System.out.println("Item added: " + item);
        vendingMachine.setState(vendingMachine.getHasItemState());
    }

    @Override
    public void addMoney(double amount) {
        System.out.println("No item available to add money for.");
    }

    @Override
    public void dispenseItem() {
        System.out.println("No item available.");
    }
}

class VendingMachine {
    private Lock lock = new ReentrantLock();
    private State hasItemState;
    private State itemRequestedState;
    private State hasMoneyState;
    private State noItemState;
    private State state;

    public VendingMachine() {
        hasItemState = new HasItemState(this);
        itemRequestedState = new ItemRequestedState(this);
        hasMoneyState = new HasMoneyState(this);
        noItemState = new NoItemState(this);
        state = noItemState;
    }

    public void setState(State state) {
        lock.lock();
        try {
            this.state = state;
        } finally {
            lock.unlock();
        }
    }

    public void selectItem() {
        lock.lock();
        try {
            state.selectItem();
        } finally {
            lock.unlock();
        }
    }

    public void addItem(String item) {
        lock.lock();
        try {
            state.addItem(item);
        } finally {
            lock.unlock();
        }
    }

    public void addMoney(double amount) {
        lock.lock();
        try {
            if (amount <= 0) {
                System.err.println("Invalid amount. Must be greater than zero.");
                return;
            }
            state.addMoney(amount);
        } finally {
            lock.unlock();
        }
    }

    public void dispenseItem() {
        lock.lock();
        try {
            state.dispenseItem();
        } finally {
            lock.unlock();
        }
    }

    public State getHasItemState() {
        return hasItemState;
    }

    public State getItemRequestedState() {
        return itemRequestedState;
    }

    public State getHasMoneyState() {
        return hasMoneyState;
    }

    public State getNoItemState() {
        return noItemState;
    }
}

public class StateDemo {
    public static void main(String[] args) {
        VendingMachine vendingMachine = new VendingMachine();

        vendingMachine.addItem("Soda");
        vendingMachine.selectItem();
        vendingMachine.addMoney(1.0);
        vendingMachine.dispenseItem();
    }
}
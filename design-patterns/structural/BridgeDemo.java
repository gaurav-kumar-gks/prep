/**
 * Bridge is a structural design pattern that lets you split a large class or a
 * set of closely related classes
 * into two separate hierarchies—abstraction and implementation—which can be
 * developed independently of each other.
 */

// Implementation interface
interface Device {
    void turnOn();

    void turnOff();

    void setVolume(int volume);

    void setChannel(int channel);
}

// Concrete implementations
class TV implements Device {
    private boolean isOn = false;
    private int volume = 30;
    private int channel = 1;

    @Override
    public void turnOn() {
        isOn = true;
        System.out.println("TV is turned on");
    }

    @Override
    public void turnOff() {
        isOn = false;
        System.out.println("TV is turned off");
    }

    @Override
    public void setVolume(int volume) {
        this.volume = volume;
        System.out.println("TV volume set to " + volume);
    }

    @Override
    public void setChannel(int channel) {
        this.channel = channel;
        System.out.println("TV channel set to " + channel);
    }
}

class Radio implements Device {
    private boolean isOn = false;
    private int volume = 30;
    private int channel = 1;

    @Override
    public void turnOn() {
        isOn = true;
        System.out.println("Radio is turned on");
    }

    @Override
    public void turnOff() {
        isOn = false;
        System.out.println("Radio is turned off");
    }

    @Override
    public void setVolume(int volume) {
        this.volume = volume;
        System.out.println("Radio volume set to " + volume);
    }

    @Override
    public void setChannel(int channel) {
        this.channel = channel;
        System.out.println("Radio channel set to " + channel);
    }
}

// Abstraction
abstract class RemoteControl {
    protected Device device;

    public RemoteControl(Device device) {
        this.device = device;
    }

    public abstract void turnOn();

    public abstract void turnOff();

    public abstract void volumeUp();

    public abstract void volumeDown();

    public abstract void channelUp();

    public abstract void channelDown();
}

// Refined abstraction
class BasicRemote extends RemoteControl {
    public BasicRemote(Device device) {
        super(device);
    }

    @Override
    public void turnOn() {
        device.turnOn();
    }

    @Override
    public void turnOff() {
        device.turnOff();
    }

    @Override
    public void volumeUp() {
        device.setVolume(device instanceof TV ? 50 : 40);
    }

    @Override
    public void volumeDown() {
        device.setVolume(device instanceof TV ? 10 : 20);
    }

    @Override
    public void channelUp() {
        device.setChannel(device instanceof TV ? 10 : 5);
    }

    @Override
    public void channelDown() {
        device.setChannel(device instanceof TV ? 1 : 1);
    }
}

public class BridgeDemo {
    public static void main(String[] args) {
        System.out.println("Testing TV with basic remote:");
        Device tv = new TV();
        RemoteControl tvRemote = new BasicRemote(tv);
        tvRemote.turnOn();
        tvRemote.volumeUp();
        tvRemote.channelUp();
        tvRemote.turnOff();

        System.out.println("\nTesting Radio with basic remote:");
        Device radio = new Radio();
        RemoteControl radioRemote = new BasicRemote(radio);
        radioRemote.turnOn();
        radioRemote.volumeDown();
        radioRemote.channelDown();
        radioRemote.turnOff();
    }
}
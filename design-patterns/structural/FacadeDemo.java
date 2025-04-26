/**
 * Facade is a structural design pattern that provides a simplified interface to
 * a complex subsystem.
 * It defines a higher-level interface that makes the subsystem easier to use.
 */

// Complex subsystem classes
class CPU {
    public void freeze() {
        System.out.println("CPU: Freezing...");
    }

    public void jump(long position) {
        System.out.println("CPU: Jumping to position " + position);
    }

    public void execute() {
        System.out.println("CPU: Executing...");
    }
}

class Memory {
    public void load(long position, byte[] data) {
        System.out.println("Memory: Loading data at position " + position);
    }
}

class HardDrive {
    public byte[] read(long lba, int size) {
        System.out.println("HardDrive: Reading " + size + " bytes from LBA " + lba);
        return new byte[size];
    }
}

// Facade
class ComputerFacade {
    private CPU cpu;
    private Memory memory;
    private HardDrive hardDrive;

    public ComputerFacade() {
        this.cpu = new CPU();
        this.memory = new Memory();
        this.hardDrive = new HardDrive();
    }

    public void start() {
        System.out.println("Computer: Starting...");
        cpu.freeze();
        memory.load(0, hardDrive.read(0, 1024));
        cpu.jump(0);
        cpu.execute();
        System.out.println("Computer: Started successfully");
    }

    public void shutdown() {
        System.out.println("Computer: Shutting down...");
        // Add shutdown logic here
        System.out.println("Computer: Shutdown complete");
    }
}

// Client
public class FacadeDemo {
    public static void main(String[] args) {
        ComputerFacade computer = new ComputerFacade();

        System.out.println("Starting computer:");
        computer.start();

        System.out.println("\nShutting down computer:");
        computer.shutdown();
    }
}
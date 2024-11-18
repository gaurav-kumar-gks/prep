from abc import ABC, abstractmethod


class OTP(ABC):

    def generate_and_send_otp(self):
        self.generate_otp()
        self.send_notification(self.get_message())

    @abstractmethod
    def generate_otp(self):
        pass

    @abstractmethod
    def get_message(self) -> str:
        pass

    @abstractmethod
    def send_notification(self, string: str):
        pass


class SMSOTP(OTP):

    def generate_otp(self):
        print("Generating SMS OTP")

    def get_message(self):
        return "SMS OTP"

    def send_notification(self, message: str):
        print(f"Sending SMS with Message: {message}")


class EmailOTP(OTP):

    def generate_otp(self):
        print("Generating Email OTP")

    def get_message(self):
        return "Email OTP"

    def send_notification(self, message: str):
        print(f"Sending Email with Message: {message}")


if __name__ == "__main__":
    smsOTP = SMSOTP()
    smsOTP.generate_and_send_otp()

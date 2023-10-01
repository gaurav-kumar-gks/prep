class KafkaProducer:
    __instance = None

    @staticmethod
    def get_producer(*args, **kwargs):
        """
        Static access method.
        """

        if KafkaProducer.__instance is None:
            KafkaProducer.__instance = KafkaProducer(*args, **kwargs)

        return KafkaProducer.__instance

    def __init__(self, conf, **kwargs):
        """
            Virtually private constructor.
        """
        self.kafka_config = conf
        self.kafka_topic = kwargs.get('topic', '')
        print("KafkaProducer init")
        if KafkaProducer.__instance:
            raise Exception("This class is a singleton!")
        else:
            try:
                self.kafka_producer = 'Producer(self.kafka_config)'
                print(f"[KafkaProducer] Connected to kafka producer successfully")
            except Exception as e:
                print(f"[KafkaProducer] Error in kafka producer init: {e} with kafka config: {self.kafka_config}")
                raise e


print("start")
conf = {}
kfp1 = KafkaProducer.get_producer(conf)
kfp2 = KafkaProducer.get_producer(conf)
# kfp1 = KafkaProducer(conf)
# kfp2 = KafkaProducer(conf)
print("Done")

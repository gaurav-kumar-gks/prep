import pandas as pd

df = pd.DataFrame(
    {
        "name": ["John", "Smith", "Paul"],
        "age": [25, 30, 35],
        "city": ["New York", "Los Angeles", "San Francisco"],
    }
)

df.head(2)
df.tail(2)
df.describe()
df.info()
print(df.shape)
print(df.dtypes)

# print(df)  # dataframe
# print(df["name"])  # series
# print(df[['name', 'age']])  # name and age
# print(df[df["age"] > 30])  # where age > 30
# print(df[df["city"].str.contains("New")])  # where city contains "New"
# print(df[df["city"].isin(["New York", "Los Angeles"])])  # where city is in list
# print(df[df["city"].notnull()])  # where city is not null
# print(df.loc[1])  # row 1
print(df.loc[1, "name"])  # row 1, column name
df.iloc()
ages = pd.Series([25, 30, 35], name="age")
ages.max()
print(ages.shape)

# df = pd.read_csv("filename.csv")
# df = pd.read_excel()
# df.to_csv("filename.csv")
# df.to_excel()





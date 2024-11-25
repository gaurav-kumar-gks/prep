import pandas as pd
import numpy as np

# Read data from a CSV file
df = pd.read_csv('data.csv')

df.head()  # Output: First 5 rows of the DataFrame
df.tail()  # Output: Last 5 rows of the DataFrame
df.info()  # Output: Summary of the DataFrame including data types and non-null counts
df.describe()  # Output: Summary statistics for numerical columns

column_name = df.columns[0]  # Assume the first column exists
df[column_name]  # Output: Series of the first column
selected_columns = df[['col1', 'col2']]  # Output: DataFrame with 'col1' and 'col2'
filtered_df = df[df[column_name] > 0]  # Output: DataFrame with rows where the first column's value is greater than 0

df['new_column'] = np.random.rand(len(df))
df.head()  # Output: DataFrame with a new column 'new_column' added

df_dropped = df.drop(columns=['new_column'])
df_dropped.head()  # Output: DataFrame after dropping 'new_column'

grouped = df.groupby(column_name).mean() # Output: DataFrame with mean values grouped by the first column

df_filled = df.fillna(0) # Handle missing data
df_filled.head()  # Output: DataFrame with missing values filled with 0

df2 = pd.DataFrame({'key': [1, 2, 3], 'value': ['a', 'b', 'c']})
merged_df = pd.merge(df, df2, left_on=column_name, right_on='key', how='inner')
# Output: Merged DataFrame based on the 'key' column

# Pivot table
pivot_table = df.pivot_table(values=column_name, index='col1', columns='col2', aggfunc='mean')
# Output: Pivot table with mean values

# Save DataFrame to a new CSV file
df.to_csv('output.csv', index=False)
# Output: DataFrame saved to 'output.csv'

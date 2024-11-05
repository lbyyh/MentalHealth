import sys
import json
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.preprocessing import LabelEncoder, StandardScaler
from sklearn.pipeline import Pipeline
from sklearn.metrics import accuracy_score

# 加载数据
df = pd.read_csv(r'D:\code\GO\MentalHealth-Platform\app\pymodel\college_survey.csv')

# 对分类变量进行编码
label_encoders = {}
for column in ['grade', 'gender', 'birthplace', 'monthly_expense']:
    label_encoders[column] = LabelEncoder().fit(df[column])
    df[column] = label_encoders[column].transform(df[column])

# 划分特征和目标
features = df[['home_town', 'expectations', 'single_child', 'future_job_expectation', 'relationship_with_classmates', 'exam_tasks', 'ability_to_handle', 'care_about_others', 'self_requirement', 'impact_by_grade', 'impact_by_gender']]
target = df['level']

# 划分训练集和测试集
features_train, features_test, target_train, target_test = train_test_split(features, target, test_size=0.2, random_state=42)

# 创建并训练模型
model = Pipeline([('scaler', StandardScaler()), ('classifier', RandomForestClassifier(random_state=42))])
model.fit(features_train, target_train)

# 从命令行参数获取输入值来进行预测
json_str = sys.argv[1]
data = json.loads(json_str)

input_data = {k: int(v) for k, v in data.items()}

def predict(input_data):
    return model.predict(pd.DataFrame(input_data, index=[0]))

# 预测并打印结果
print(predict(input_data))

# 验证模型效果
predictions = model.predict(features_test)
accuracy = accuracy_score(target_test, predictions)
print(f'Accuracy: {accuracy}')
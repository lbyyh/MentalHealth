#!/usr/bin/env python
import sys

def predict_mental_health_status(name, age, gender, mood, work_stress, social_activity, sleep_quality):
    """
    一个简单的心理健康状态预测函数。
    这个函数根据输入的参数返回一个预测结果。
    实际应用中，这里应该有一个复杂的机器学习模型。
    """
    # 这里只是一个示例，实际的预测逻辑会更复杂
    if age < 18:
        return "未成年，建议寻求家长或监护人的帮助。"

    if work_stress == "high" and sleep_quality == "low":
        return "工作压力大且睡眠质量差，可能存在心理健康风险。"

    if mood == "low" and social_activity == "low":
        return "情绪低落且社交活动少，建议寻求专业心理咨询。"

    return "目前没有明显的心理健康风险。"

if __name__ == "__main__":
    # 解析命令行参数
    if len(sys.argv) != 8:
        print("Usage: python predict.py <name> <age> <gender> <mood> <work_stress> <social_activity> <sleep_quality>")
        sys.exit(1)

    name, age, gender, mood, work_stress, social_activity, sleep_quality = sys.argv[1:]

    # 将年龄从字符串转换为整数
    try:
        age = int(age)
    except ValueError:
        print("年龄必须是数字。")
        sys.exit(1)

    # 调用预测函数并打印结果
    result = predict_mental_health_status(name, age, gender, mood, work_stress, social_activity, sleep_quality)
    print(result)
@[title:title]
# 2장 파이썬의 변수와 기본자료형

# 학습목표
1. 변수를 선언하고 데이터를 저장하는 방법을 배우세요.
2. 파이썬의 기본 자료형(숫자형, 문자열, 불리언)을 익히세요.
3. 자료형 간 연산과 변환을 실습하세요.
4. 파이썬의 자료형과 변수 활용법을 통해 간단한 문제를 해결하세요.


## 1. 변수와 데이터 저장
프로그래밍에서 변수(variable) 는 데이터를 저장하는 하나의 공간(space)을 말합니다. 변수를 사용하면 데이터를 쉽게 저장하고 재사용할 수 있습니다. Python에서 변수를 이해하고 올바르게 사용하는 방법을 배워봅시다.

---

### **1.1 변수란 무엇인가?**

변수는 데이터를 저장할 수 있는 **이름이 있는 공간**입니다. 마치 **라벨이 붙은 상자**처럼, 변수에 값을 저장하고 필요할 때 불러올 수 있습니다.

@[b003:example]

```python
x = 5  # 변수 x에 5 저장
print(x + 3)  # 결과: 8
```

---


### **1.2 파이썬의 특징: 동적 타이핑 (Dynamic Typing)**

파이썬은 변수를 선언할 때 자료형(int, str 등)을 미리 정하지 않음. 변수에 값이 할당되는 시점에 실시간으로 자료형이 결정되는 방식을 **동적 타이핑**이라고 함.

- **특징:** 하나의 변수에 처음에 숫자를 넣었다가, 나중에 문자열을 다시 넣는 것이 가능함.
- **장점:** 코드가 간결하고 유연함.

>>>runnable
```python
# 1. 정수 저장
data = 10
print(data, type(data))  # 결과: 10 <class 'int'>

# 2. 동일한 변수에 문자열 재할당
data = "Hello Python"
print(data, type(data))  # 결과: Hello Python <class 'str'>

# 3. 동일한 변수에 실수 재할당
data = 3.14
print(data, type(data))  # 결과: 3.14 <class 'float'>
```

> 💡 **참고:** `type()` 함수는 해당 변수가 현재 어떤 자료형을 가지고 있는지 확인하는 데 사용함.
> 

---

@[runnable]
+++ toml:data
language = "python"
version = "3.12"
dependencies = []
code = """
```python
# 1. 정수 저장
data = 10
print(data, type(data))  # 결과: 10 <class 'int'>

# 2. 동일한 변수에 문자열 재할당
data = "Hello Python"
print(data, type(data))  # 결과: Hello Python <class 'str'>

# 3. 동일한 변수에 실수 재할당
data = 3.14
print(data, type(data))  # 결과: 3.14 <class 'float'>
```
"""
+++

@[choice]
+++ toml:data
prompt = "blah blah"
options = ["aaa","bbb","ccc"]
answers = [0]

+++

@[collapsible]
+++ toml:data
summary = "click to expand"
+++ md:body
This content will be hidden by default. 
You can put any other Markdown formatting here, such as:

*   Lists
*   **Bold** text
*   Code blocks
+++

@[short_answer]
+++ toml:data
prompt = "blah blah"
answer = "answer"
+++
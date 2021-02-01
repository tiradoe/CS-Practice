def removeDupes(numbers):
    print("Starting Length: ", len(numbers))
    for i in range(len(numbers) - 1,0,-1):
        if numbers[i] == numbers[i - 1]:
            numbers.pop(i)

    print(numbers)

    return len(numbers)

numbers = [0,0,1,1,1,2,2,3,3]
deduped = removeDupes(numbers)
print("Final Length: ", deduped)


def removeDuplicates(nums) -> int:
    print(nums)
    print("Len: ", len(nums))
    i = 0
    for number in nums:
        print("number: ", number)
        print("current: ", nums[i])
        if number + 1 == len(nums):
            continue

        if nums[i] == nums[i+1]:
            nums.remove(nums[i])
            print("Removing: ", nums[i])

        i = i + 1

    return len(nums)


length = removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4])
print(length)

--
-- @lc app=leetcode.cn id=182 lang=mysql
--
-- [182] 查找重复的电子邮箱
--

-- @lc code=start
# Write your MySQL query statement below

select Email from Person where Id In (select Id from Person Group by Email Having Count(*)>1)
-- @lc code=end


--
-- @lc app=leetcode.cn id=181 lang=mysql
--
-- [181] 超过经理收入的员工
--

-- @lc code=start
# Write your MySQL query statement below

select Name as Employee from Employee t1 where Salary > (select Salary from Employee where Id = t1.ManagerId)
-- @lc code=end


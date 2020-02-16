--
-- @lc app=leetcode.cn id=183 lang=mysql
--
-- [183] 从不订购的客户
--

-- @lc code=start
# Write your MySQL query statement below

select Name Customers from Customers  where Id not in (select CustomerId from Orders)
-- @lc code=end


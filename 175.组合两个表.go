
/*
 * @lc app=leetcode.cn id=175 lang=mysql
 *
 * [175] 组合两个表
 */

// @lc code=start

select Person.FirstName, Person.LastName, Address.City, Address.State from Person  
  left join Address  on Person.PersonId = Address.PersonId

// @lc code=end
select max(Salary) from Employee limit 1 where Salary < (select max(Salary) from Employee)

select max(Salary) SecondHighestSalary
from employee
where
salary<(select max(salary) from employee)
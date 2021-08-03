/**
 * @name GORM error not checked
 * @description A call that interact with the database using the GORM library
 *              without checking whether there was an error.
 * @kind problem
 * @problem.severity warning
 * @id turboscan/go/gorm-error-not-checked
 * @precision high
 */

import go
 
from CallExpr ce
where 
  ce.getType().(PointerType).getBaseType().getQualifiedName() = "github.com/jinzhu/gorm.DB" and
  not exists(Assignment a | ce.getParent+() = a) and
  not exists(ReturnStmt s | ce.getParent+() = s) and
  not exists(SelectorExpr se | se = ce.getParent+() | se.getSelector().getName() = "Error" ) and
  ce.getCalleeName() != "InstantSet" and
  ce.getCalleeName() != "LogMode"
select ce, "This call appears to interact with the database without checking whether an error was encountered."

package sale_stock

import (
	"github.com/hexya-erp/pool/h"
)

//vars

var (
	//Enable Route on Sales Order Line
	GroupRouteSoLines *security.Group
)


//rights
func init() {
	h.Stock.ModelStockPicking().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelStockPicking().Methods().Write().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelStockPicking().Methods().Create().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelStockMove().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelStockMove().Methods().Write().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelStockMove().Methods().Create().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelStockMove().Methods().AllowAllToGroup(GroupSaleManager)
	h.Procurement.ModelProcurementOrder().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Procurement.ModelProcurementOrder().Methods().Write().AllowGroup(GroupSaleSalesman)
	h.Procurement.ModelProcurementOrder().Methods().Create().AllowGroup(GroupSaleSalesman)
	h.Procurement.ModelProcurementOrder().Methods().AllowAllToGroup(GroupSaleManager)
	h.SaleOrder().Methods().Load().AllowGroup(GroupStockUser)
	h.SaleOrder().Methods().Write().AllowGroup(GroupStockUser)
	h.SaleOrderLine().Methods().Load().AllowGroup(GroupStockUser)
	h.SaleOrderLine().Methods().Write().AllowGroup(GroupStockUser)
	h.Stock.ModelStockPicking().Methods().AllowAllToGroup(GroupSaleManager)
	h.Product.ModelProductPackaging().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Product.ModelProductPackaging().Methods().Write().AllowGroup(GroupSaleSalesman)
	h.Product.ModelProductPackaging().Methods().Create().AllowGroup(GroupSaleSalesman)
	h.Product.ModelProductPackaging().Methods().Load().AllowGroup(GroupSaleManager)
	h.Stock.ModelStockWarehouse().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelStockLocation().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Product.ModelProductPackaging().Methods().AllowAllToGroup(GroupSaleManager)
	h.Stock.ModelStockWarehouseOrderpoint().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Account.ModelAccountPartialReconcile().Methods().AllowAllToGroup(GroupStockManager)
	h.Account.ModelAccountJournal().Methods().Load().AllowGroup(GroupStockManager)
	h.Account.ModelAccountJournal().Methods().Write().AllowGroup(GroupStockManager)
	h.Stock.ModelStockLocation().Methods().Load().AllowGroup(GroupSaleManager)
	h.Stock.ModelProcurementRule().Methods().AllowAllToGroup(GroupSaleManager)
	h.Stock.ModelStockLocationPath().Methods().AllowAllToGroup(GroupSaleManager)
	h.Stock.ModelStockLocationPath().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Stock.ModelProcurementRule().Methods().Load().AllowGroup(GroupSaleSalesman)
}

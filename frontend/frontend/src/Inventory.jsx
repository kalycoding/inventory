import './App.css';
import React from 'react';
import axios from 'axios';



function InventoryApp(){
    const [formValue, setformValue] = React.useState({
        product_id:'',
        cost_price: 0,
        quantity: 0,
        stock_level: 0,
        selling_price:0,
        supplier:''
      });  
    
      const handleDelete = (data)=> {
        axios.delete(`http://localhost:8080/api/inventory/${data}`)
              .then(function(response){
             
              })
              window.location.reload(); 
      }
    
      const [inventoryData, setInventory] = React.useState([]);
    
      React.useEffect(() => {
        let mounted = true;
        axios.get('http://localhost:8080/api/inventory')
              .then(function(response){
                if (response.status===200){
                  if(mounted) {
                      setInventory(response.data)
                    }
                }
                
              })
              .catch(error => {setInventory([])
              })
        
        // getList()
        //   .then(items => {
        //     
        //   })
        return () => mounted = false;
      }, [])
    
      console.log(inventoryData)
      const handleSubmit = (event) => {
        //event.preventDefault();
        const productFormData = new FormData();
        productFormData.append("product_id", formValue.product_id)
        productFormData.append("cost_price", formValue.cost_price)
        productFormData.append("quantity", formValue.quantity)
        productFormData.append("stock_level", formValue.stock_level)
        productFormData.append("selling_price", formValue.selling_price)
        productFormData.append("supplier", formValue.supplier)
        const body = {
            costPrice : formValue.cost_price, 
            quantity : formValue.quantity,
            stockLevel : formValue.stock_level,
            sellingPrice : formValue.selling_price,
            supplier : formValue.supplier
        };
          axios.post('http://localhost:8080/api/inventory',body, { params: { prodId: formValue.product_id } })
              .then(function(response){
                if (response.status===201){
                  alert(`Category ${formValue.supplier} Created`)
                }
                else{
                  alert(`Category ${formValue.supplier} Already Exist`)
                }
              })
              .catch(error => {alert(error)
              })
          }
          const handleChange = (event) => {
            setformValue({
              ...formValue,
              [event.target.name]: event.target.value
            });
          }

          return(
            <div>
            <form onSubmit={handleSubmit}>
              <p>Inventory Creation Page</p>
              <input
                type="text"
                name="product_id"
                placeholder="Product"
                value={formValue.product_id}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="cost_price"
                placeholder="Cost Price"
                value={formValue.cost_price}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="quantity"
                placeholder="Quantity"
                value={formValue.quantity}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="stock_level"
                placeholder="Stock Level"
                value={formValue.stock_level}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="selling_price"
                placeholder="Selling Price"
                value={formValue.selling_price}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="supplier"
                placeholder="Supplier"
                value={formValue.supplier}
                onChange={handleChange}
                required
              />
              <button
                type="submit"
              >
                Create Inventory
              </button>
              <p>Inventory List</p> 
            </form>
            {
            !inventoryData ? <p>No data</p> : 
            <ul>
                {inventoryData.map(data=>(<button key={data.id} 
                  onClick={()=>handleDelete(data.id)}>
                    {data.product.name}</button>)
                )}
            </ul>
            }
            </div>
          )
}

export default InventoryApp
import React from 'react';
import axios from 'axios';
import './Category.css';
function InventoryApp(){
    const [formValue, setformValue] = React.useState({
        product_id:'',
        cost_price: 0,
        quantity: 0,
        stock_level: 0,
        selling_price:0,
        supplier:'',
        edit_cost_price: '',
        edit_quantity: '',
        edit_stock_level: '',
        edit_selling_price:'',
        edit_supplier:''
      });  
    
      const handleDelete = (data)=> {
        axios.delete(`http://localhost:8080/api/inventory/${data}`)
              .then(function(response){
             
              })
              window.location.reload(); 
      }
    
    
      const [inventoryData, setInventory] = React.useState([]);
      const [productData, setProduct] = React.useState([]);
      
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
      React.useEffect(() => {
        let mounted = true;
        axios.get('http://localhost:8080/api/product')
              .then(function(response){
                if (response.status===200){
                  if(mounted) {
                      setProduct(response.data)
                    }
                }
                
              })
              .catch(error => {setProduct([])
              })
        
        // getList()
        //   .then(items => {
        //     
        //   })
        return () => mounted = false;
      }, [])
    
      console.log(formValue)
      const handleSubmit = (event) => {
        //event.preventDefault();
        const productFormData = new FormData();
        productFormData.append("products", formValue.product_id)
        productFormData.append("cost_price", formValue.cost_price)
        productFormData.append("quantity", formValue.quantity)
        productFormData.append("stock_level", formValue.stock_level)
        productFormData.append("selling_price", formValue.selling_price)
        productFormData.append("supplier", formValue.supplier)
        productFormData.append("edit_cost_price", formValue.edit_cost_price)
        productFormData.append("edit_quantity", formValue.edit_quantity)
        productFormData.append("edit_stock_level", formValue.edit_stock_level)
        productFormData.append("edit_selling_price", formValue.edit_selling_price)
        productFormData.append("edit_supplier", formValue.supplier)
        const body = {
            costPrice : Number(formValue.cost_price), 
            quantity : Number(formValue.quantity),
            stockLevel : Number(formValue.stock_level),
            sellingPrice : Number(formValue.selling_price),
            supplier : formValue.supplier
        };
          axios.post('http://localhost:8080/api/inventory',body, { params: { prodId: formValue.product_id } })
              .then(function(response){
                // if (response.status===201){
                //   alert(`Category ${formValue.supplier} Created`)
                // }
                // else{
                //   alert(`Category ${formValue.supplier} Already Exist`)
                // }
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
            <h1 class="title">Inventory</h1>
            <p>
            {
            !inventoryData ? <p></p> : <ul>
              <p><center><button class="export">Export Inventory to CSV</button></center></p>
              </ul>
              
          } 
            </p>
        <div class="container">
          <form onSubmit={handleSubmit}>
          
            <div class="row">
              <div class="col-25">
                <label for="fname">Product</label>
              </div>
              <div class="col-75">
                {
                  !productData ? <select><option value="none" selected disabled hidden >Empty Product, Add Product</option></select> : 
                  <select name="product_id" onChange={handleChange}>
                    <option value="none" selected disabled hidden>Choose Category</option>
                    {productData.map(data=>(
                    <option key={data.id} value={data.id} name="product_id">{data.name}</option>
                    ))}
                  </select>
                }
              </div>
              <div class="col-25">
                <label for="fname">Cost Price</label>
              </div>
              <div class="col-75">
                <input
                  type="text"
                  name="cost_price"
                  placeholder="Cost Price"
                  value={formValue.cost_price}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>
            <div class="row">
              <div class="col-25">
                <label for="fname">Quantity</label>
              </div>
              <div class="col-75">
                <input
                  type="text"
                  name="quantity"
                  placeholder="Quantity"
                  value={formValue.quantity}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>
            <div class="row">
              <div class="col-25">
                <label for="fname">Stock Level</label>
              </div>
              <div class="col-75">
                <input
                  type="text"
                  name="stock_level"
                  placeholder="Stock Level"
                  value={formValue.stock_level}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>
            <div class="row">
              <div class="col-25">
                <label for="fname">Selling Price</label>
              </div>
              <div class="col-75">
                <input
                  type="text"
                  name="selling_price"
                  placeholder="Selling Price"
                  value={formValue.selling_price}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>
            <div class="row">
              <div class="col-25">
                <label for="fname">Supplier</label>
              </div>
              <div class="col-75">
                <input
                  type="text"
                  name="supplier"
                  placeholder="Supplier"
                  value={formValue.supplier}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>
            <div class="row">
              <input type="submit" value="Submit"/>
            </div>
          </form>
        </div>  
        <div>
        <p>
          
          {
            !inventoryData ? <p>Empty Inventory, Please Add</p> : <ul>
              {inventoryData.map(data=>(<span>
              <form>
              <li key={data.id}>
                <table>
                  <tr>
                    <th>Product</th>
                    <th>Quantity</th>
                    <th>Cost Price</th>
                    <th>Stock Level</th>
                    <th>Selling Price</th>
                    <th>Supplier</th>
                  </tr>
                  <tr>
                    <td>{data.product.name}</td>
                    <td>
                    <input
                      type="text"
                      name="edit_quantity"
                      placeholder={data.quantity}
                      value={data.quantity}
                      onChange={handleChange}
                      required
                    />
                    </td>
                    <td>
                      <input
                        type="text"
                        name="edit_cost_price"
                        placeholder="Cost Price"
                        value={data.costPrice}
                        onChange={handleChange}
                        required
                      />
                    </td>
                    <td>
                      <input
                        type="text"
                        name="edit_stock_level"
                        placeholder="Stock Level"
                        value={data.stockLevel}
                        onChange={handleChange}
                        required
                      />
                    </td>
                    <td>
                      <input
                        type="text"
                        name="edit_selling_price"
                        placeholder="Selling Price"
                        value={data.sellingPrice}
                        onChange={handleChange}
                        required
                      />
                    </td>
                    <td>
                      <input
                        type="text"
                        name="edit_supplier"
                        placeholder="Supplier"
                        value={data.supplier}
                        onChange={handleChange}
                        required
                      />
                    </td>
                  </tr>
              </table>
              </li>
              
              <p><button class="edit">Edit</button></p>
              </form>
              <p><button class="delete" onClick={()=>handleDelete(data.id)}>Delete</button></p>
              </span>))}</ul>
              
          } 
        </p>
      </div>
    </div>
        ) 
}

export default InventoryApp
/* <div>
            <form onSubmit={handleSubmit} id='products'>
              <p>Inventory Creation Page</p>
              

                {
                  !productData ? <select><option value="none" selected disabled hidden>Empty Products, Add product</option></select> : 
                  <select name="product_id" onChange={handleChange}>
                    <option value="none" selected disabled hidden>Choose Product</option>
                    {productData.map(data=>(
                    <option key={data.id} value={data.id} name="product_id">{data.name}</option>
                    
                    ))}
                  </select>
                }
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
            </div> */
const mysql = require('mysql2/promise')
const md5 = require("crypto-js/md5");

const pool = mysql.createPool({
  host: process.env.DATABASE_HOST,
  port: process.env.DATABASE_PORT,
  user: 'dev',
  password: 'qV2qe9nXZjQ0c4CT',
  database: 'm1337Shirts',
  waitForConnections: true,
  connectionLimit: 10,
  maxIdle: 10, 
  idleTimeout: 60000,
  queueLimit: 0,
  enableKeepAlive: true,
  keepAliveInitialDelay: 0,
});

const gunPool = mysql.createPool({
  host: process.env.DATABASE_HOST,
  port: process.env.DATABASE_PORT,
  user: 'root',
  password: 'BHBU2p2ChyY9ctVxKyIxlZ4gqL5tpQrg',
  database: 'm1337Guns',
  waitForConnections: true,
  connectionLimit: 10,
  maxIdle: 10, 
  idleTimeout: 60000,
  queueLimit: 0,
  enableKeepAlive: true,
  keepAliveInitialDelay: 0,
});

async function getProducts(max_photos=null, search=null, guns=false) {
  if (typeof(search) != 'string') search = ""
  else search = search.replace("!", "!!").replace("%", "!%").replace("_", "!_").replace("[", "![");
  search = '%' + search + '%'

  let products
  let conn
  const prodPool = guns ? gunPool : pool

  try {
    let conn = await prodPool.getConnection();
    
    let photos;

    let [productsRes, _] = await conn.query(products_search_query, [search])
    const productDict = productsRes.map(p => p.id).reduce((acc, key, index) => {
      acc[key] = productsRes[index];
      acc[key].photos = []
      return acc;
    }, {});

    if (max_photos != null) {
      const [rows, _] = await conn.query(images_search_maxImages_query, [search, max_photos])
      photos = rows
    } else {
      const [rows, _] = await conn.query(images_search_query, [search])
      photos = rows
    }

    photos.forEach(photo => {
      if (productDict[photo.productId]) {
        productDict[photo.productId].photos.push(photo.filepath)
      }
    })

    products = Object.values(productDict)
  } catch (err) {
    console.log(err)
    products = []
  } finally {
    if (conn) {
      prodPool.releaseConnection(conn);
    }
  }

  return products
}

async function getProduct(product_id, max_photos=null, guns=true) {
  
  let product;
  let conn
  const prodPool = guns ? gunPool : pool

  try {
    conn = await prodPool.getConnection();
    
    let photos = [];

    const [productsRes, _] = await conn.execute(product_byId_query, [product_id])
    if (productsRes.length == 0) {
      product = {}
    } else {
      product = productsRes[0]
      product.photos = []
  
      if (max_photos != null) {
        const [rows, _] = await conn.execute(images_byId_maxImages_query, [product_id, max_photos])
        photos = rows
      } else {
        const [rows, _] = await conn.execute(images_byId_query, [product_id])
        photos = rows
      }
      
      photos.forEach(photo => {
        product.photos.push(photo.filepath)
      })
    }

  } catch (err) {
    console.log(err)
    product = null
  } finally {
    if (conn) {
      prodPool.releaseConnection(conn);
    }
  }

  return product
}

async function getUserId(username, password) {
  
  let userId
  let conn

  try {
    conn = await pool.getConnection();

    const hashed_password = md5(password).toString()
    let query = `SELECT id FROM user_profile WHERE username = '${username}' and hashed_password = '${hashed_password}' LIMIT 1;`
    
    const [rows, _] = await conn.query(query, [])

    pool.releaseConnection(conn);

    userId = rows?.[0]?.id ?? -1
  } catch (err) {
    userId = null
  } finally {
    if (conn) {
      pool.releaseConnection(conn);
    }
  }

  return userId
}

async function getUser(id) {
  
  let user
  let conn

  try {  
    conn = await pool.getConnection();
    const [rows, _] = await conn.query(user_byId_query, [id])

    user = rows?.[0] ?? {}
  } catch (err) {
    console.log(err)
    user = {}
  } finally {
    pool.releaseConnection(conn);
  }

  return user
}

async function getPromo(code) {  
  let promo
  let conn

  try {
    conn = await pool.getConnection();
    const [rows, _] = await conn.query(promo_byCode_query, [code])

    promo = rows?.[0]
  } catch (err) {
    console.log(err)
    promo = {}
  } finally {
    if (conn) {
      pool.releaseConnection(conn)
    }
  }

  return promo
}

module.exports = {
  getProducts, getProduct, getUserId, getUser, getPromo
}

const products_search_query = `SELECT * FROM product WHERE title LIKE ? ESCAPE '!' ORDER BY id;`
const product_byId_query = `SELECT * FROM product WHERE id = ? ORDER BY id;`
const user_byId_query = `SELECT * FROM user_profile WHERE id = ? LIMIT 1;`
const promo_byCode_query = `SELECT * FROM promo where promo_code = ?;`

const images_search_query = `SELECT productId, filepath from product_image
INNER JOIN (
	SELECT id FROM product
  WHERE title LIKE ? ESCAPE '!'
) product_searched ON product_searched.id = product_image.productId
ORDER BY productId, displayOrder;`

const images_search_maxImages_query = `SELECT productId, filepath from product_image
INNER JOIN (
	SELECT id FROM product
  WHERE title LIKE ? ESCAPE '!'
) product_searched ON product_searched.id = product_image.productId
WHERE displayOrder < ?
ORDER BY productId, displayOrder;`

const images_byId_maxImages_query = `SELECT productId, filepath FROM product_image
INNER JOIN product ON product.id = product_image.productId
WHERE productId = ? and displayOrder < ?
ORDER BY productId, displayOrder;`

const images_byId_query = `SELECT productId, filepath FROM product_image
INNER JOIN product ON product.id = product_image.productId
WHERE productId = ?
ORDER BY productId, displayOrder;`
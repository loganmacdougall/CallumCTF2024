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

  const conn = await (guns ? gunPool.getConnection() : pool.getConnection());

  try {
    let photos;

    const [products, _] = await conn.query(products_search_query, [search])
    const productDict = products.map(p => p.id).reduce((acc, key, index) => {
      acc[key] = products[index];
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

    pool.releaseConnection(conn);

    return Object.values(productDict)
  } catch (err) {

    pool.releaseConnection(conn);
    console.log(err)

    return []
  }
}

async function getProduct(product_id, max_photos=null, guns=true) {
  const conn = await (guns ? gunPool.getConnection() : pool.getConnection());

  try {
    let photos = [];

    const [products, _] = await conn.execute(product_byId_query, [product_id])
    if (products.length == 0) {
      return {}
    }
    const product = products[0]
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

    pool.releaseConnection(conn);

    return product
  } catch (err) {
    pool.releaseConnection(conn);
    console.log(err)

    return null
  }
}

async function getUserId(username, password) {
  
  const conn = await pool.getConnection();
  
  try {
    const hashed_password = md5(password).toString()
    let query = `SELECT id FROM user_profile WHERE username = '${username}' and hashed_password = '${hashed_password}' LIMIT 1;`
    
    const [rows, _] = await conn.query(query, [])

    pool.releaseConnection(conn);

    return rows?.[0]?.id ?? -1
  } catch (err) {
    pool.releaseConnection(conn);
    
    return null
  }
}

async function getUser(id) {
  const conn = await pool.getConnection();
  
  try {  
    const [rows, _] = await conn.query(user_byId_query, [id])

    pool.releaseConnection(conn);

    return rows?.[0] ?? {}
  } catch (err) {
    console.log(err)
    pool.releaseConnection(conn);
    
    return {}
  }
}

async function getPromo(code) {
  const conn = await pool.getConnection();
  let promo = {}

  try {
    const [rows, _] = await conn.query(promo_byCode_query, [code])

    promo = rows?.[0]

  } catch (err) {
    console.log(err)
  }

  pool.releaseConnection(conn)
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
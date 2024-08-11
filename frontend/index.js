const express = require('express')
const cookieParser = require('cookie-parser')
const session = require('express-session')
const conn = require('./database')

const app = express()

const port = 3000

const global_state = {promo: true}

app.set('view engine', 'ejs');
app.set('trust proxy', 1) // trust first proxy
app.set('views', './views')

app.use('/static', express.static('public'))
app.use(express.urlencoded({extended: true}))
app.use(express.json())
app.use(cookieParser())
app.use(session({
  secret: 'keyboard dog',
  resave: false,
  saveUninitialized: true,
  cookie: { secure: false }
}))

app.get('/', (req, res) => {
  if (seized(res)) return
  res.render('home', {session: req.session})
})

app.get('/shop', async (req, res) => {
  if (seized(res)) return
  let search = req.query.search
  if (search == undefined) search = null

  const products = await conn.getProducts(max_photos=1, search=search)
  res.render('shop', {session: req.session, products: products, global_state: global_state})
})

app.get('/product', async (req, res) => {
  if (seized(res)) return

  let id = req.query.id
  if (isNaN(id)) {
    return res.redirect("/")
  }

  promo = {}
  let hasDisconut = false, discount = 0
  let reveal = false

  if (global_state.promo == true && req.query.promo) {
    promo = await conn.getPromo(req.query.promo)
    if (promo) {
      let promo_action = promo.promo_action.split(" ")
      if (promo_action[0] == 'DISCOUNT') {
        hasDisconut = true
        discount = Number.parseInt(promo_action[1])
      } else if (promo_action[0] == 'REVEAL') {
        reveal = true
      }
    }
  }

  let product
  if (reveal) {
    product = {
      id: 9999,
      title: "On the Inside",
      price: 1199.99,
      pdesc: "Get access to hardware and communication channels exclusive to those on the inside",
      photos: ["/media/products/handshake.webp"]
    }
  } else {
    product = await conn.getProduct(id)
    if (product.id == undefined) {
      return res.redirect("/")
    }
    if (hasDisconut) {
      product.price = `<s>$${product.price}</s> $${Math.floor(100 * (product.price - (product.price * discount * 0.01))) / 100}`
    } else {
      product.price = `$${product.price}`
    }
  }

  res.render('product', {session: req.session, product: product, promo: promo, global_state: global_state})
})

app.post('/buynow', async (req, res) => {
  if (seized(res)) return
  if (req.session.user_id == undefined) {
    return res.sendStatus('403')
  }

  if (req.body.id == undefined) {
    return res.sendStatus('400')
  }

  if (global_state.promo && req.body.id == 9999) {
    return res.status(302).send("/inside?key=only_the_know_truly_know")
  }

  let product = await conn.getProduct(req.body.id)

  if (product == null) {
    return res.sendStatus('404')
  }

  res.sendStatus('200')
})

app.get('/login', (req, res) => {
  if (seized(res)) return
  req.session.user_id = undefined;
  req.session.user_username = undefined;

  res.render('login', {session: req.session, status: "", global_state: global_state})
})

app.post('/login', async (req, res) => {
  if (seized(res)) return
  let username = req.body.username
  let password = req.body.password

  if (username == undefined || password == undefined) {
    let message = "Didn\'t receive credentials";
    console.log(message)
    return res.status(400).send(message)
  }

  let id = await conn.getUserId(username, password)

  if (id == -1) {
    let message = "Invalid credentials";
    return res.status(403).send(message)
  } else if (id == null) {
    let message = "Server Error";
    return res.status(500).send(message)
  }

  req.session.user_id = id;
  req.session.user_username = username;

  res.sendStatus(200)
})

app.get('/profile', async (req, res) => {
  if (seized(res)) return
  if (req.session.user_id == undefined) {
    return res.redirect('/login')
  }

  const user = await conn.getUser(req.session.user_id)

  res.render('profile', {session: req.session, user: user, global_state: global_state})
})

app.get('/thankyou', async (req, res) => {
  if (seized(res)) return
  let id = req.query.id
  if (isNaN(id)) {
    return res.redirect("/")
  }

  let product = await conn.getProduct(id)
  if (product.id == undefined) {
    return res.redirect("/")
  }

  res.render('thankyou', {session: req.session, product: product, global_state: global_state})
})

app.get('/inside', (req, res) => {
  if (seized(res) || !global_state.promo) return

  let key = req.query.key
  if (!key || key != 'only_the_know_truly_know') return

  res.render('inside', {session: req.session, global_state: global_state})
})

app.patch('/reveal_new_truths', (req, res) => {
  if (seized(res)) return
  const key = 'ruRNZcHJu59BZXuAP24N9Z4zqAN6GmUJ'
  const provided_key = req.body.key ?? ''
  const action = req.body.action ?? ''

  if (key != provided_key) {
    return
  }

  global_state[action] = true;

  res.sendStatus(200)
}) 

const server = app.listen(port,() => {
  console.log(`Server is running on http://localhost:${port}`);
});

process.on('SIGTERM', () => {
  console.log('SIGTERM signal received: closing HTTP server')
  server.close(() => {
    console.log('HTTP server closed')
  })
})

function seized(res) {
  
  if ( global_state.seized == true) {
    res.render('csci')
    return true
  }
  return false
}
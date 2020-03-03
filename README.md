# Deli

Deli platform users:

- People looking to takeout
- Partners
  - Restaurants looking to expand their clientele
  - Boda-Boda looking for a flexible and competitive income



## Requirement

#### User(takeout/takeaway user)

```go
User {
    firstname
    lastname
    email_address
    phone_number
    password
    address
}
```



#### Partners

- Restaurants

  ```go
  Restaurant {
      restaurant_name
      telephone
      delicacies{}
      address {
          postal_code
          building_name/number
          street
          country
      }
      Licenses {
          business_permit
          food_hygiene_license
          food_handler_certificate
          fire_certificate
          environment_certificate
      }
      verified
      ratings
      reviews[]
  }
  ```

- Boda-Boda riders

  ```go
  firstname
  lastname
  email_address
  phone_number
  password
  city
  legal_documents {
      driving_permit
      PSV_driving_license
      first_aid_training_certificate
  }
  ```

  

### User stories

- Sign-up/Sign-in
- allow location history
- CRUD home address
- view cuisines from nearest restaurants  if`location == true`
- view restaurants with good ratings `rating > 4.5` if `location == false`
- place order on different cuisines
- add payout method
- make payment for orders and deliveries if `payoutMethod == true`
- **Partners**
  - *Restaurants*
    - CRUD menu
    - view order(s)
    - accept order
    - track payments
    - add payout methods
  - *Boda-Boda delivery*
    - confirm delivery
    - reject delivery
    - view delivery process in real-time
    - complete meal delivery




package database

var stations_data = `[
  {
    "name": "Acton Town",
    "line": ["Piccadilly", "District"],
    "zone": [3]
  },
  {
    "name": "Aldgate",
    "line": ["Metropolitan", "Circle"],
    "zone": [1]
  },
  {
    "name": "Aldgate East",
    "line": ["Hammersmith and City", "District"],
    "zone": [1]
  },
  {
    "name": "Alperton",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Amersham",
    "line": ["Metropolitan"],
    "zone": [9]
  },
  {
    "name": "Angel",
    "line": ["Northern"],
    "zone": [1]
  },
  {
    "name": "Archway",
    "line": ["Northern"],
    "zone": [2, 3]
  },
  {
    "name": "Arnos Grove",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Arsenal",
    "line": ["Piccadilly"],
    "zone": [2]
  },
  {
    "name": "Baker Street",
    "line": [
      "Metropolitan",
      "Bakerloo",
      "Circle",
      "Jubilee",
      "Hammersmith and City"
    ],
    "zone": [1]
  },
  {
    "name": "Balham",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "Bank",
    "line": ["Waterloo and City", "Northern", "Central", "DLR"],
    "zone": [1]
  },
  {
    "name": "Barbican",
    "line": ["Metropolitan", "Circle", "Hammersmith and City"],
    "zone": [1]
  },
  {
    "name": "Barking",
    "line": ["District", "Hammersmith and City"],
    "zone": [4]
  },
  {
    "name": "Barkingside",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Barons Court",
    "line": ["District", "Piccadilly"],
    "zone": [2]
  },
  {
    "name": "Battersea Power Station",
    "line": ["Northern"],
    "zone": [1]
  },
  {
    "name": "Bayswater",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Becontree",
    "line": ["District"],
    "zone": [5]
  },
  {
    "name": "Belsize Park",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Bermondsey",
    "line": ["Jubilee"],
    "zone": [2]
  },
  {
    "name": "Bethnal Green",
    "line": ["Central"],
    "zone": [2]
  },
  {
    "name": "Blackfriars",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Blackhorse Road",
    "line": ["Victoria"],
    "zone": [3]
  },

  {
    "name": "Bond Street",
    "line": ["Central", "Jubilee"],
    "zone": [1]
  },
  {
    "name": "Borough",
    "line": ["Northern"],
    "zone": [1]
  },
  {
    "name": "Boston Manor",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Bounds Green",
    "line": ["Piccadilly"],
    "zone": [3, 4]
  },
  {
    "name": "Bow Road",
    "line": ["District", "Hammersmith and City"],
    "zone": [2]
  },
  {
    "name": "Brent Cross",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "Brixton",
    "line": ["Victoria"],
    "zone": [2]
  },
  {
    "name": "Bromley-by-Bow",
    "line": ["District", "Hammersmith and City"],
    "zone": [2, 3]
  },
  {
    "name": "Buckhurst Hill",
    "line": ["Central"],
    "zone": [5]
  },
  {
    "name": "Burnt Oak",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "Caledonian Road",
    "line": ["Piccadilly"],
    "zone": [2]
  },
  {
    "name": "Camden Town",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Canada Water",
    "line": ["Jubilee"],
    "zone": [2]
  },
  {
    "name": "Canary Wharf",
    "line": ["Jubilee", "DLR"],
    "zone": [2]
  },
  {
    "name": "Canning Town",
    "line": ["Jubilee", "DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Cannon Street",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Canons Park",
    "line": ["Jubilee"],
    "zone": [5]
  },
  {
    "name": "Chalfont and Latimer",
    "line": ["Metropolitan"],
    "zone": [8]
  },
  {
    "name": "Chalk Farm",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Chancery Lane",
    "line": ["Central"],
    "zone": [1]
  },
  {
    "name": "Charing Cross",
    "line": ["Bakerloo", "Northern"],
    "zone": [1]
  },
  {
    "name": "Chesham",
    "line": ["Metropolitan"],
    "zone": [9]
  },
  {
    "name": "Chigwell",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Chiswick Park",
    "line": ["District"],
    "zone": [3]
  },
  {
    "name": "Chorleywood",
    "line": ["Metropolitan"],
    "zone": [7]
  },
  {
    "name": "Clapham Common",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Clapham North",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Clapham South",
    "line": ["Northern"],
    "zone": [2, 3]
  },
  {
    "name": "Cockfosters",
    "line": ["Piccadilly"],
    "zone": [5]
  },
  {
    "name": "Colindale",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "Colliers Wood",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "Covent Garden",
    "line": ["Piccadilly"],
    "zone": [1]
  },
  {
    "name": "Croxley",
    "line": ["Metropolitan"],
    "zone": [7]
  },
  {
    "name": "Dagenham East",
    "line": ["District"],
    "zone": [5]
  },
  {
    "name": "Dagenham Heathway",
    "line": ["District"],
    "zone": [5]
  },
  {
    "name": "Debden",
    "line": ["Central"],
    "zone": [6]
  },
  {
    "name": "Dollis Hill",
    "line": ["Jubilee"],
    "zone": [3]
  },
  {
    "name": "Ealing Broadway",
    "line": ["District", "Central", "Elizabeth"],
    "zone": [3]
  },
  {
    "name": "Ealing Common",
    "line": ["District", "Piccadilly"],
    "zone": [3]
  },
  {
    "name": "Earl's Court",
    "line": ["District", "Piccadilly"],
    "zone": [1, 2]
  },
  {
    "name": "East Acton",
    "line": ["Central"],
    "zone": [2]
  },
  {
    "name": "East Finchley",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "East Ham",
    "line": ["District", "Hammersmith and City"],
    "zone": [3, 4]
  },
  {
    "name": "East Putney",
    "line": ["District"],
    "zone": [2, 3]
  },
  {
    "name": "Eastcote",
    "line": ["Metropolitan", "Piccadilly"],
    "zone": [5]
  },
  {
    "name": "Edgware",
    "line": ["Northern"],
    "zone": [5]
  },
  {
    "name": "Edgware Road",
    "line": ["Bakerloo"],
    "zone": [1]
  },
  {
    "name": "Elephant and Castle",
    "line": ["Northern", "Bakerloo"],
    "zone": [1, 2]
  },
  {
    "name": "Elm Park",
    "line": ["District"],
    "zone": [6]
  },
  {
    "name": "Embankment",
    "line": ["District", "Bakerloo", "Northern", "Circle"],
    "zone": [1]
  },
  {
    "name": "Epping",
    "line": ["Central"],
    "zone": [6]
  },
  {
    "name": "Euston",
    "line": ["Northern", "Victoria"],
    "zone": [1]
  },
  {
    "name": "Euston Square",
    "line": ["Metropolitan", "Circle", "Hammersmith and City"],
    "zone": [1]
  },
  {
    "name": "Fairlop",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Farringdon",
    "line": ["Metropolitan", "Circle", "Hammersmith and City"],
    "zone": [1]
  },
  {
    "name": "Finchley Central",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "Finchley Road",
    "line": ["Metropolitan", "Jubilee"],
    "zone": [2]
  },
  {
    "name": "Finsbury Park",
    "line": ["Piccadilly", "Victoria"],
    "zone": [2]
  },
  {
    "name": "Fulham Broadway",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "Gants Hill",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Gloucester Road",
    "line": ["District", "Piccadilly", "Circle"],
    "zone": [1]
  },
  {
    "name": "Golders Green",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "Goldhawk Road",
    "line": ["Hammersmith and City", "Circle"],
    "zone": [2]
  },
  {
    "name": "Goodge Street",
    "line": ["Northern"],
    "zone": [1]
  },
  {
    "name": "Grange Hill",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Great Portland Street",
    "line": ["Metropolitan", "Circle", "Hammersmith and City"],
    "zone": [1]
  },
  {
    "name": "Greenford",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Green Park",
    "line": ["Piccadilly", "Victoria", "Jubilee"],
    "zone": [1]
  },
  {
    "name": "Gunnersbury",
    "line": ["District"],
    "zone": [3]
  },
  {
    "name": "Hainault",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Hammersmith",
    "line": ["District", "Piccadilly"],
    "zone": [2]
  },
  {
    "name": "Hampstead",
    "line": ["Northern"],
    "zone": [2, 3]
  },
  {
    "name": "Hanger Lane",
    "line": ["Central"],
    "zone": [3]
  },
  {
    "name": "Harlesden",
    "line": ["Bakerloo"],
    "zone": [3]
  },
  {
    "name": "Harrow and Wealdstone",
    "line": ["Bakerloo"],
    "zone": [5]
  },
  {
    "name": "Harrow-on-the-Hill",
    "line": ["Metropolitan"],
    "zone": [5]
  },
  {
    "name": "Hatton Cross",
    "line": ["Piccadilly"],
    "zone": [5, 6]
  },
  {
    "name": "Heathrow Terminals 2 and 3",
    "line": ["Piccadilly", "Elizabeth"],
    "zone": [6]
  },
  {
    "name": "Heathrow Terminal 4",
    "line": ["Piccadilly", "Elizabeth"],
    "zone": [6]
  },
  {
    "name": "Heathrow Terminal 5",
    "line": ["Piccadilly", "Elizabeth"],
    "zone": [6]
  },
  {
    "name": "Hendon Central",
    "line": ["Northern"],
    "zone": [3, 4]
  },
  {
    "name": "High Barnet",
    "line": ["Northern"],
    "zone": [5]
  },
  {
    "name": "Highbury and Islington",
    "line": ["Victoria"],
    "zone": [2]
  },
  {
    "name": "Highgate",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "High Street Kensington",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Hillingdon",
    "line": ["Metropolitan", "Piccadilly"],
    "zone": [6]
  },
  {
    "name": "Holborn",
    "line": ["Central", "Piccadilly"],
    "zone": [1]
  },
  {
    "name": "Holland Park",
    "line": ["Central"],
    "zone": [2]
  },
  {
    "name": "Holloway Road",
    "line": ["Piccadilly"],
    "zone": [2]
  },
  {
    "name": "Hornchurch",
    "line": ["District"],
    "zone": [6]
  },
  {
    "name": "Hounslow Central",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Hounslow East",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Hounslow West",
    "line": ["Piccadilly"],
    "zone": [5]
  },
  {
    "name": "Hyde Park Corner",
    "line": ["Piccadilly"],
    "zone": [1]
  },
  {
    "name": "Ickenham",
    "line": ["Metropolitan", "Piccadilly"],
    "zone": [6]
  },
  {
    "name": "Kennington",
    "line": ["Northern"],
    "zone": [1, 2]
  },
  {
    "name": "Kensal Green",
    "line": ["Bakerloo"],
    "zone": [2]
  },
  {
    "name": "Kensington (Olympia)",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "Kentish Town",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Kenton",
    "line": ["Bakerloo"],
    "zone": [4]
  },
  {
    "name": "Kew Gardens",
    "line": ["District"],
    "zone": [3, 4]
  },
  {
    "name": "Kilburn",
    "line": ["Jubilee"],
    "zone": [2]
  },
  {
    "name": "Kilburn Park",
    "line": ["Bakerloo"],
    "zone": [2]
  },
  {
    "name": "Kingsbury",
    "line": ["Jubilee"],
    "zone": [4]
  },
  {
    "name": "King's Cross St Pancras",
    "line": [
      "Metropolitan",
      "Northern",
      "Piccadilly",
      "Circle",
      "Victoria",
      "Hammersmith and City"
    ],
    "zone": [1]
  },
  {
    "name": "Knightsbridge",
    "line": ["Piccadilly"],
    "zone": [1]
  },
  {
    "name": "Ladbroke Grove",
    "line": ["Hammersmith and City", "Circle"],
    "zone": [2]
  },
  {
    "name": "Lambeth North",
    "line": ["Bakerloo"],
    "zone": [1]
  },
  {
    "name": "Lancaster Gate",
    "line": ["Central"],
    "zone": [1]
  },
  {
    "name": "Latimer Road",
    "line": ["Hammersmith and City", "Circle"],
    "zone": [2]
  },
  {
    "name": "Leicester Square",
    "line": ["Piccadilly", "Northern"],
    "zone": [1]
  },
  {
    "name": "Leyton",
    "line": ["Central"],
    "zone": [3]
  },
  {
    "name": "Leytonstone",
    "line": ["Central"],
    "zone": [3, 4]
  },
  {
    "name": "Liverpool Street",
    "line": ["Metropolitan", "Central", "Circle", "Hammersmith and City"],
    "zone": [1]
  },
  {
    "name": "London Bridge",
    "line": ["Northern", "Jubilee"],
    "zone": [1]
  },
  {
    "name": "Loughton",
    "line": ["Central"],
    "zone": [6]
  },
  {
    "name": "Maida Vale",
    "line": ["Bakerloo"],
    "zone": [2]
  },
  {
    "name": "Manor House",
    "line": ["Piccadilly"],
    "zone": [2, 3]
  },
  {
    "name": "Mansion House",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Marble Arch",
    "line": ["Central"],
    "zone": [1]
  },
  {
    "name": "Marylebone",
    "line": ["Bakerloo"],
    "zone": [1]
  },
  {
    "name": "Mile End",
    "line": ["District", "Hammersmith and City", "Central"],
    "zone": [2]
  },
  {
    "name": "Mill Hill East",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "Monument",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Moorgate",
    "line": ["Northern", "Circle", "Hammersmith and City"],
    "zone": [1]
  },
  {
    "name": "Moor Park",
    "line": ["Metropolitan"],
    "zone": [6, 7]
  },
  {
    "name": "Morden",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "Mornington Crescent",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Neasden",
    "line": ["Jubilee"],
    "zone": [3]
  },
  {
    "name": "Newbury Park",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Nine Elms",
    "line": ["Northern"],
    "zone": [1]
  },
  {
    "name": "North Acton",
    "line": ["Central"],
    "zone": [2, 3]
  },
  {
    "name": "North Ealing",
    "line": ["Piccadilly"],
    "zone": [3]
  },
  {
    "name": "North Greenwich",
    "line": ["Jubilee"],
    "zone": [2, 3]
  },
  {
    "name": "North Harrow",
    "line": ["Metropolitan"],
    "zone": [5]
  },
  {
    "name": "North Wembley",
    "line": ["Bakerloo"],
    "zone": [4]
  },
  {
    "name": "Northfields",
    "line": ["Piccadilly"],
    "zone": [3]
  },
  {
    "name": "Northolt",
    "line": ["Central"],
    "zone": [5]
  },
  {
    "name": "Northwick Park",
    "line": ["Metropolitan"],
    "zone": [4]
  },
  {
    "name": "Northwood",
    "line": ["Metropolitan"],
    "zone": [6]
  },
  {
    "name": "Northwood Hills",
    "line": ["Metropolitan"],
    "zone": [6]
  },
  {
    "name": "Notting Hill Gate",
    "line": ["District", "Central", "Circle"],
    "zone": [1, 2]
  },
  {
    "name": "Oakwood",
    "line": ["Piccadilly"],
    "zone": [5]
  },
  {
    "name": "Old Street",
    "line": ["Northern"],
    "zone": [1]
  },
  {
    "name": "Osterley",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Oval",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Oxford Circus",
    "line": ["Central", "Bakerloo", "Victoria"],
    "zone": [1]
  },
  {
    "name": "Paddington",
    "line": ["District", "Circle", "Bakerloo"],
    "zone": [1]
  },
  {
    "name": "Park Royal",
    "line": ["Piccadilly"],
    "zone": [3]
  },
  {
    "name": "Parsons Green",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "Perivale",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Piccadilly Circus",
    "line": ["Bakerloo", "Piccadilly"],
    "zone": [1]
  },
  {
    "name": "Pimlico",
    "line": ["Victoria"],
    "zone": [1]
  },
  {
    "name": "Pinner",
    "line": ["Metropolitan"],
    "zone": [5]
  },
  {
    "name": "Plaistow",
    "line": ["District", "Hammersmith and City"],
    "zone": [3]
  },
  {
    "name": "Preston Road",
    "line": ["Metropolitan"],
    "zone": [4]
  },
  {
    "name": "Putney Bridge",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "Queen's Park",
    "line": ["Bakerloo"],
    "zone": [2]
  },
  {
    "name": "Queensbury",
    "line": ["Jubilee"],
    "zone": [4]
  },
  {
    "name": "Queensway",
    "line": ["Central"],
    "zone": [1]
  },
  {
    "name": "Ravenscourt Park",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "Rayners Lane",
    "line": ["Metropolitan", "Piccadilly"],
    "zone": [5]
  },
  {
    "name": "Redbridge",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Regent's Park",
    "line": ["Bakerloo"],
    "zone": [1]
  },
  {
    "name": "Richmond",
    "line": ["District"],
    "zone": [4]
  },
  {
    "name": "Rickmansworth",
    "line": ["Metropolitan"],
    "zone": [7]
  },
  {
    "name": "Roding Valley",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Royal Oak",
    "line": ["Hammersmith and City", "Circle"],
    "zone": [2]
  },
  {
    "name": "Ruislip",
    "line": ["Metropolitan", "Piccadilly"],
    "zone": [6]
  },
  {
    "name": "Ruislip Gardens",
    "line": ["Central"],
    "zone": [5]
  },
  {
    "name": "Ruislip Manor",
    "line": ["Metropolitan", "Piccadilly"],
    "zone": [6]
  },
  {
    "name": "Russell Square",
    "line": ["Piccadilly"],
    "zone": [1]
  },
  {
    "name": "St James's Park",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "St John's Wood",
    "line": ["Jubilee"],
    "zone": [2]
  },
  {
    "name": "St Paul's",
    "line": ["Central"],
    "zone": [1]
  },
  {
    "name": "Seven Sisters",
    "line": ["Victoria"],
    "zone": [3]
  },
  {
    "name": "Shepherd's Bush",
    "line": ["Central"],
    "zone": [2]
  },
  {
    "name": "Shepherd's Bush Market",
    "line": ["Hammersmith and City", "Circle"],
    "zone": [2]
  },
  {
    "name": "Sloane Square",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Snaresbrook",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "South Ealing",
    "line": ["Piccadilly"],
    "zone": [3]
  },
  {
    "name": "South Harrow",
    "line": ["Piccadilly"],
    "zone": [5]
  },
  {
    "name": "South Kensington",
    "line": ["District", "Piccadilly", "Circle"],
    "zone": [1]
  },
  {
    "name": "South Kenton",
    "line": ["Bakerloo"],
    "zone": [4]
  },
  {
    "name": "South Ruislip",
    "line": ["Central"],
    "zone": [5]
  },
  {
    "name": "South Wimbledon",
    "line": ["Northern"],
    "zone": [3, 4]
  },
  {
    "name": "South Woodford",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Southfields",
    "line": ["District"],
    "zone": [3]
  },
  {
    "name": "Southgate",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Southwark",
    "line": ["Jubilee"],
    "zone": [1]
  },
  {
    "name": "Stamford Brook",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "Stanmore",
    "line": ["Jubilee"],
    "zone": [5]
  },
  {
    "name": "Stepney Green",
    "line": ["District", "Hammersmith and City"],
    "zone": [2]
  },
  {
    "name": "Stockwell",
    "line": ["Northern", "Victoria"],
    "zone": [2]
  },
  {
    "name": "Stonebridge Park",
    "line": ["Bakerloo"],
    "zone": [3]
  },
  {
    "name": "Stratford",
    "line": ["Central", "Jubilee", "DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Sudbury Hill",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Sudbury Town",
    "line": ["Piccadilly"],
    "zone": [4]
  },
  {
    "name": "Swiss Cottage",
    "line": ["Jubilee"],
    "zone": [2]
  },
  {
    "name": "Temple",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Theydon Bois",
    "line": ["Central"],
    "zone": [6]
  },
  {
    "name": "Tooting Bec",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "Tooting Broadway",
    "line": ["Northern"],
    "zone": [3]
  },
  {
    "name": "Tottenham Court Road",
    "line": ["Central", "Northern"],
    "zone": [1]
  },
  {
    "name": "Tottenham Hale",
    "line": ["Victoria"],
    "zone": [3]
  },
  {
    "name": "Totteridge and Whetstone",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "Tower Hill",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Tufnell Park",
    "line": ["Northern"],
    "zone": [2]
  },
  {
    "name": "Turnham Green",
    "line": ["District", "Piccadilly"],
    "zone": [2, 3]
  },
  {
    "name": "Turnpike Lane",
    "line": ["Piccadilly"],
    "zone": [3]
  },
  {
    "name": "Upminster",
    "line": ["District"],
    "zone": [6]
  },
  {
    "name": "Upminster Bridge",
    "line": ["District"],
    "zone": [6]
  },
  {
    "name": "Upney",
    "line": ["District"],
    "zone": [4]
  },
  {
    "name": "Upton Park",
    "line": ["District", "Hammersmith and City"],
    "zone": [3]
  },
  {
    "name": "Uxbridge",
    "line": ["Metropolitan", "Piccadilly"],
    "zone": [6]
  },
  {
    "name": "Vauxhall",
    "line": ["Victoria"],
    "zone": [1, 2]
  },
  {
    "name": "Victoria",
    "line": ["District", "Circle"],
    "zone": [1]
  },
  {
    "name": "Walthamstow Central",
    "line": ["Victoria"],
    "zone": [3]
  },
  {
    "name": "Wanstead",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Warren Street",
    "line": ["Northern", "Victoria"],
    "zone": [1]
  },
  {
    "name": "Warwick Avenue",
    "line": ["Bakerloo"],
    "zone": [2]
  },
  {
    "name": "Waterloo",
    "line": ["Waterloo and City", "Bakerloo", "Northern", "Jubilee"],
    "zone": [1]
  },
  {
    "name": "Watford",
    "line": ["Metropolitan"],
    "zone": [7]
  },
  {
    "name": "Wembley Central",
    "line": ["Bakerloo"],
    "zone": [4]
  },
  {
    "name": "Wembley Park",
    "line": ["Metropolitan", "Jubilee"],
    "zone": [4]
  },
  {
    "name": "West Acton",
    "line": ["Central"],
    "zone": [3]
  },
  {
    "name": "West Brompton",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "West Finchley",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "West Ham",
    "line": ["District", "Hammersmith and City", "Jubilee", "DLR"],
    "zone": [2, 3]
  },
  {
    "name": "West Hampstead",
    "line": ["Jubilee"],
    "zone": [2]
  },
  {
    "name": "West Harrow",
    "line": ["Metropolitan"],
    "zone": [5]
  },
  {
    "name": "West Kensington",
    "line": ["District"],
    "zone": [2]
  },
  {
    "name": "West Ruislip",
    "line": ["Central"],
    "zone": [6]
  },
  {
    "name": "Westbourne Park",
    "line": ["Hammersmith and City", "Circle"],
    "zone": [2]
  },
  {
    "name": "Westminster",
    "line": ["District", "Circle", "Jubilee"],
    "zone": [1]
  },
  {
    "name": "White City",
    "line": ["Central"],
    "zone": [2]
  },
  {
    "name": "Whitechapel",
    "line": ["District", "Hammersmith and City"],
    "zone": [2]
  },
  {
    "name": "Willesden Green",
    "line": ["Jubilee"],
    "zone": [2, 3]
  },
  {
    "name": "Willesden Junction",
    "line": ["Bakerloo"],
    "zone": [2, 3]
  },
  {
    "name": "Wimbledon",
    "line": ["District"],
    "zone": [3]
  },
  {
    "name": "Wimbledon Park",
    "line": ["District"],
    "zone": [3]
  },
  {
    "name": "Wood Green",
    "line": ["Piccadilly"],
    "zone": [3]
  },
  {
    "name": "Wood Lane",
    "line": ["Hammersmith and City", "Circle"],
    "zone": [2]
  },
  {
    "name": "Woodford",
    "line": ["Central"],
    "zone": [4]
  },
  {
    "name": "Woodside Park",
    "line": ["Northern"],
    "zone": [4]
  },
  {
    "name": "Abbey Road",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "All Saints",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Beckton",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Beckton Park",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Blackwall",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Bow Church",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Crossharbour",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Custom House",
    "line": ["Elizabeth", "DLR"],
    "zone": [3]
  },
  {
    "name": "Cutty Sark",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Cyprus",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Deptford Bridge",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Devons Road",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "East India",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Elverson Road",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Gallions Reach",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Greenwich National Rail",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Heron Quays",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Island Gardens",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "King George V",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Langdon Park",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Lewisham",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Limehouse",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "London City Airport",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Mudchute",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Pontoon Dock",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Poplar",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Prince Regent",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Pudding Mill Lane",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Royal Albert",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Royal Victoria",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Shadwell",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "South Quay",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Star Lane",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Stratford High Street",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Stratford International",
    "line": ["DLR"],
    "zone": [2, 3]
  },
  {
    "name": "Tower Gateway",
    "line": ["DLR"],
    "zone": [1]
  },
  {
    "name": "West India Quay",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "West Silvertown",
    "line": ["DLR"],
    "zone": [3]
  },
  {
    "name": "Westferry",
    "line": ["DLR"],
    "zone": [2]
  },
  {
    "name": "Woolwich Arsenal",
    "line": ["DLR"],
    "zone": [4]
  },
  {
    "name": "West Drayton",
    "line": ["Elizabeth"],
    "zone": [6]
  },
  {
    "name": "Hayes and Harlington",
    "line": ["Elizabeth"],
    "zone": [5]
  },
  {
    "name": "Southall",
    "line": ["Elizabeth"],
    "zone": [4]
  },
  {
    "name": "Hanwell",
    "line": ["Elizabeth"],
    "zone": [4]
  },
  {
    "name": "West Ealing",
    "line": ["Elizabeth"],
    "zone": [3]
  },
  {
    "name": "Acton Main Line",
    "line": ["Elizabeth"],
    "zone": [3]
  },
  {
    "name": "Abbey Wood",
    "line": ["Elizabeth"],
    "zone": [4]
  },
  {
    "name": "Maryland",
    "line": ["Elizabeth"],
    "zone": [3]
  },
  {
    "name": "Forest Gate",
    "line": ["Elizabeth"],
    "zone": [3]
  },
  {
    "name": "Manor Park",
    "line": ["Elizabeth"],
    "zone": [3, 4]
  },
  {
    "name": "Ilford",
    "line": ["Elizabeth"],
    "zone": [4]
  },
  {
    "name": "Seven Kings",
    "line": ["Elizabeth"],
    "zone": [4]
  },
  {
    "name": "Goodmayes",
    "line": ["Elizabeth"],
    "zone": [4]
  },
  {
    "name": "Chadwell Heath",
    "line": ["Elizabeth"],
    "zone": [4]
  },
  {
    "name": "Romford",
    "line": ["Elizabeth"],
    "zone": [6]
  },
  {
    "name": "Gidea Park",
    "line": ["Elizabeth"],
    "zone": [6]
  },
  {
    "name": "Harold Wood",
    "line": ["Elizabeth"],
    "zone": [6]
  },
  {
    "name": "Brentwood",
    "line": ["Elizabeth"],
    "zone": [9]
  }
]`

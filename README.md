# Composable Data Service GraphQL Proof-of-Concept

## How to run the service locally
### Requirements
 - Go

Run the command below:
```bash
go run server.go
```

## Overview
This is a proof of concept that showcases how the Composable Data Service (CDS) is built using GraphQL to handle the graph nature of the problem the service is built to solve.

## Problem
We need a service that enables clients to query data in a shape that could be considered complex.

```
- get me markets related to "generative ai"
    - for each of those markets, I want:
        - basic market data (name, industry, subindustry, etc.)
        - extended market data (description, key features, who buys it, who it serves)
        - top 10 companies in the market, each with:
            - basic company data (name, location, URLs)
            - extended company data (description, headcount, revenue, etc.)
            - top 3 competitor companies, each with:
                - basic company data only
        - top 3 related markets, each with:
            - basic market data
            - top 5 companies in the market, each with:
                - basic company data only
```

Given this potential request shown above, it can be tricky to fetch the data needed and resolve them in a shape the client wants. This example is pulled from [here](https://miro.com/app/board/uXjVLtguFgw=/). You'll also notice that there is data that could be fetched in a batch rather than fetching individually. For example, multiple fields pertaining to different (or potentially the same) companies. A graph can be realized with the relationships between markets and companies as well as between different companies.

## Potential Solution
GraphQL is an excellent choice as an API for a service that should provide data in a graph-like nature for the following reasons:

### 1. Graph Representation of Data
 - GraphQL inherently models data as a graph, where entities and relationships can be queried in a structured and efficient way.
 - This aligns well with services dealing with interconnected data. From the perspective of CDS, the relationships between markets and companies.
   - What companies are in a certain market?
   - What markets does a particular company belong to?
   - Who are the competitors for a particular company?
   - How do the funding rounds for one company compare to its competitors?

### 2. Efficient Data Retrieval
 - Traditional APIs often struggle with over-fetching (retrieving unnecessary data) or under-fetching (requiring multiple requests to assemble needed data).
 - GraphQL allows clients to specify exactly what they need, making it ideal for navigating and fetching what they need.
 - The nature of GraphQL allows to clients to reduce the number of API requests they need to make as well as omit data from responses they don't need. 
   - For example, a client will be able to, if needed, request both a list of markets and a list of companies in one request.

### 3. Flexible Querying and Traversal
 - GraphQL's query language allows clients to traverse relationships dynamically based on the client's request.
 - This prevents having to limit the traversal (recursion) level in the service. Consider the hypothetical query below. While a client might not realistically request comapnies in this shape, this query is valid and the traversal level limit will always be based on the client's request by default.

 ```graphql
query getCompanies {
  companiesByIdOrg(ids: [1,3]) {
    orgID
    name
    status
	investments {
      receiver {
        investments {
          receiver {
            investments {
              receiver {
                investments {
                  roundName
                }
              }
            }
          }
        }
      }
    }
  }
}
 ```

### 4. Optimized Data Loading with Dataloaders
 - The N+1 Query problem is very common for graph-like nature retrievals. While not specific to GraphQL, [DataLoaders](https://github.com/graphql/dataloader) are commonly used in GraphQL APIs to **efficiently batch requests** to alleviate issues regarding N+1 as long as reliable batch fetchers are available.
 - DataLoaders not only batch multiple requests but also **deduplicate** identical requests within the same request cycle. If dataloaders were used for the hypothetical query above, it will allow the resolvers for the fields in the deeper levels of the query to reuse data that was already previously fetched when available.

## Drawbacks
 - GraphQL is not currently supported in the company (At least at the time of this writing)
   - It will take some time to setup the service to behave like our other existing service. (Deployment, Security, etc.)
 - We currently lack support for data retrievals that we need for batch requests. While GraphQL can still be used to elegantly compose the data clients need, we will still be using non-batch data retrievals. 
   - It would be ideal to get support from the teams that own the data we need to fetch.
 - GraphQL will be great when paired with event-driven data system such that we have a service that collects and hydrates the data needed for Composable Data Service, which is an architecture that is yet to be realized in the company.

 ## Notable Service Dependencies
 - [gqlgen](https://github.com/99designs/gqlgen) - The Go GraphQL framework used for this POC.
 - [dataloadgen](https://github.com/vikstrous/dataloadgen) - The DataLoaders framework used for this POC.
   - This DataLoaders framework happens to not handle deduplication of data retrievals despite being recommended for most perfomant. A separate context cache has been implemented in this service to handle deduplication per request. Will continue to test other frameworks.

## GraphQL Core Concepts
 1. **Schema** - Defines the structure of the API, including types, queries, mutations. (Similar to proto in gRPC)
 2. **Type System** - The set of types to define the API's data model.
 3. **Root Types** - The entry points to a schema, including Query and Mutation.
 3. **Query** - A read operation used to fetch data from the API.
 4. **Mutation** - A write operation used to modify data on the server.
 5. **Field** - The unit of data requested from a GraphQL API; corresponds to an attribute of a type.
 6. **Resolver** - A function that fetches data for a particular field in a GraphQL query or mutation.
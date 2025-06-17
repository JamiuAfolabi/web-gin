# Sales Analytics Technical Test

## Overview
This technical test evaluates your ability to design and implement a sales analytics system using modern technologies and architectural patterns. The test focuses on your understanding of data processing, API design, and domain-driven architecture.

## Technical Requirements

### Core Technologies
- Go (Golang)
- DuckDB for data processing and analytics
- API Framework (Choose one):
  - Gin
  - Echo
  - gRPC

### Architecture Requirements
1. **Domain-Driven Design**
   - Implement clean architecture principles
   - Separate business logic from infrastructure concerns
   - Use proper domain modeling for sales and analytics

2. **Data Layer**
   - Create a mock database of sales data
   - Implement DuckDB integration for analytical queries
   - Design appropriate data models and schemas

3. **API Layer**
   - Implement RESTful endpoints or gRPC services
   - Create separate endpoints for Finance and Sales teams
   - Include proper error handling and validation

## Functional Requirements

### Data Requirements
Create a mock sales database with the following attributes:
- Product information (ID, name, category, price)
- Sales transactions (ID, product ID, quantity, date, customer info)
- Category information
- Any additional relevant fields you deem necessary

### API Endpoints
Implement the following analytical endpoints:

#### For Finance Team
1. Total sales by category
2. Revenue trends over time
3. Average order value by category
4. Top performing products
5. Sales performance metrics

#### For Sales Team
1. Category-wise sales distribution
2. Product performance within categories
3. Sales trends by time period
4. Customer purchase patterns
5. Category growth metrics

## Technical Evaluation Criteria

1. **Code Quality**
   - Clean, maintainable code
   - Proper error handling
   - Comprehensive documentation
   - Unit tests

2. **Architecture**
   - Proper separation of concerns
   - Domain-driven design implementation
   - Scalable and maintainable structure

3. **Performance**
   - Efficient DuckDB queries
   - Optimized data processing
   - Response time considerations

4. **API Design**
   - RESTful/gRPC best practices
   - Clear endpoint documentation
   - Proper request/response handling

## Submission Requirements

1. Source code in a Git repository
2. README with:
   - Setup instructions
   - API documentation
   - Architecture overview
   - How to run the application
3. A brief write-up (max 500 words) covering:
   - Design decisions and trade-offs
   - Potential improvements
   - Scalability considerations
   - Security considerations

## Time Frame
- Recommended completion time: 3 days
- Please indicate if you need more time

## Getting Started

1. Fork this repository
2. Set up your development environment
3. Implement the requirements
4. Document your solution
5. Submit your repository link

## Notes
- Focus on code quality over feature completeness
- Document any assumptions made
- Include any relevant diagrams
- Consider edge cases and error scenarios

## Evaluation
Your submission will be evaluated based on:
- Code quality and organization
- Architecture implementation
- API design and documentation
- Performance considerations
- Documentation quality
- Improvement suggestions

Good luck with the test! We look forward to reviewing your solution. 
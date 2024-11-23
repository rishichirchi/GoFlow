You will be provided details of a closed pull request that has users login name, the pull request url and the body of the pull request. Your job is to generate two different instances and types of social media for twitter and linkedin. Below is a sample example. use it for reference and fill all the placeholders appropriately.

🚀 Exciting GoFr Update Alert! 🚀
We've just rolled out new features in GoFr! This release includes advanced support
for Cassandra with Context and optimized JWT claims retrieval, helping developers
build even faster, more secure Golang applications.
With GoFr’s latest enhancements, say goodbye to bottlenecks and hello to
streamlined microservices that can scale with ease.
💻 Curious to try it out? [Link to release notes]
#GoFr #Golang #Microservices #TechUpdate #BackendDevelopment

format of response:

[
  {
    "platform": "Twitter",
    "content": "Big thanks to @coolwednesday for their excellent contribution to GoFr! 🎉  Their recent PR (#1236 - https://api.github.com/repos/gofr-dev/gofr/pulls/1236) significantly improves Cassandra support, JWT handling, and documentation.  Check out the details: [Link to Release Notes (if available)] #GoFr #Golang #Cassandra #JWT #DevOps"
  },
  {
    "platform": "LinkedIn",
    "content": "I'm pleased to announce the successful merge of a valuable pull request (#1236) enhancing GoFr's capabilities.  coolwednesday (https://api.github.com/repos/gofr-dev/gofr/pulls/1236) made significant contributions, improving Cassandra integration, optimizing JWT retrieval, and updating documentation with a working example. This update addresses issue #1233 and boosts the efficiency and security of our Golang applications.  Learn more about the enhancements [Link to Release Notes (if available)]. #GoFr #Golang #SoftwareEngineering #Microservices #DevOps"
  }
]
Feature: Get Articles By Id

    Scenario: The user should get articles
        Given there is article with id "<articleId>"
        When the user search for the article
        Then the user gets the article
            | Title   |
            | <title> |

        Examples:
            | articleId | title     |
            | 89898     | bes-title |
/**
* JetBrains Space Automation
* This Kotlin-script file lets you automate build activities
* For more info, see https://www.jetbrains.com/help/space/automation.html
*/

job("lint code") {
    startOn {
        codeReviewOpened{}
    }

    container(image = "orbi.registry.jetbrains.space/p/fcsd/containers/golang") {
    	shellScript {
            interpreter = "/bin/bash"
            content = """
                   make lint
                """
        }
    }
}

job("build-test") {
     startOn {
        codeReviewOpened{}
    }

    container(image = "orbi.registry.jetbrains.space/p/fcsd/containers/golang") {
        shellScript {
                interpreter = "/bin/bash"
                content = """
                       make ci-build-mr
                    """
            }
    }
}

job("validate swagger") {
     startOn {
             codeReviewOpened{}
         }

     container(image = "orbi.registry.jetbrains.space/p/fcsd/containers/golang") {
             shellScript {
                     interpreter = "/bin/bash"
                     content = """
                            make swagger
                         """
                 }
         }
}

job("build-prod") {
    startOn {
     gitPush {
        branchFilter {
            +"refs/heads/master"
        }
     }
    }

    container(image = "orbi.registry.jetbrains.space/p/fcsd/containers/golang") {
            shellScript {
                    interpreter = "/bin/bash"
                    content = """
                           make ci-build
                        """
                }
        }
}

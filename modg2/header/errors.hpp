#include <string>
#include "log.hpp"
#include "colors.hpp"

#define EXIT_WITH_ERROR(reason, program_file, logger_header, errors_header) do { \
    std::cout << std::endl << slqcmmxu << " E | " << eajvmrqh << "[-----------> " << prptbirt << " LOGGER Module ~> " << slqcmmxu << logger_header << std::endl;\
    std::cout << std::endl << slqcmmxu << " E | " << eajvmrqh << "[-----------> " << prptbirt << " LOGGER Module ~> " << slqcmmxu << Time() << std::endl;\ 
	std::cout << std::endl << slqcmmxu << " E | " << gapfwmmi << "[-----------" << slqcmmxu << "=>" << "" << " ERRORS Module ~> " << gapfwmmi << reason << std::endl;\
	exit(1); \
	} while(0)

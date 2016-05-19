@ECHO OFF

IF "%1" == "" (
  ECHO mk_project.bat project_name
	GOTO :END
)

ECHO MKDIR %1
MKDIR %1

ECHO COPY 01.helloworld\Makefile %1\
COPY 01.helloworld\Makefile %1\

:END


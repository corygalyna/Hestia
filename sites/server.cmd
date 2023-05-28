echo >/dev/null # >nul & GOTO WINDOWS & rem ^




################################################################################
# Unix Main Codes                                                              #
################################################################################
hugo server --noBuildLock \
        --disableFastRender \
        --port 8080 \
        --renderToDisk \
        --gc
exit $?
################################################################################
# Unix Main Codes                                                              #
################################################################################




:WINDOWS
::##############################################################################
:: Windows Main Codes                                                          #
::##############################################################################
hugo server --noBuildLock ^
        --disableFastRender ^
        --port 8080 ^
        --gc
::##############################################################################
:: Windows Main Codes                                                          #
::##############################################################################

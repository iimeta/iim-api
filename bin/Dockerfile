FROM loads/alpine:3.8

###############################################################################
#                                 INSTALLATION                                #
###############################################################################

ENV WORKDIR        /app
ADD bin/iim-api $WORKDIR/bin/iim-api
ADD log            $WORKDIR/log
ADD manifest       $WORKDIR/manifest
RUN chmod +x       $WORKDIR/bin/iim-api

###############################################################################
#                                    START                                    #
###############################################################################

WORKDIR $WORKDIR
CMD ./bin/iim-api

module Main exposing (..)

import Browser
import Css exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (..)
import Html.Styled.Events exposing (onClick, onInput)
import Http
import Json.Decode as D
import Json.Encode as E



-- MAIN


main =
    Browser.element
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view >> toUnstyled
        }



-- MODEL


type Model
    = Init String
    | Loading
    | Failure String
    | Success String String


init : () -> ( Model, Cmd Msg )
init _ =
    ( Init "", Cmd.none )



-- UPDATE


type Msg
    = Change String
    | Short String
    | Response (Result Http.Error String)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        Change url ->
            case model of
                Success old _ ->
                    ( Success old url, Cmd.none )

                _ ->
                    ( Init url, Cmd.none )

        Short url ->
            ( Loading, shortUrl url )

        Response result ->
            case result of
                Ok url ->
                    ( Success url "", Cmd.none )

                Err _ ->
                    ( Failure "", Cmd.none )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


theme :
    { main : Color
    , primaryLight : Color
    , secondaryLight : Color
    , primaryDark : Color
    , secondaryDark : Color
    }
theme =
    { main = hex "343477"
    , primaryLight = hex "8080B3"
    , secondaryLight = hex "565695"
    , primaryDark = hex "1A1A59"
    , secondaryDark = hex "09093B"
    }


mainContainer : List (Attribute msg) -> List (Html msg) -> Html msg
mainContainer =
    styled div
        [ displayFlex
        , alignItems center
        , border3 (px 5) solid theme.primaryDark
        , position relative
        , flexDirection column
        , justifyContent center
        , margin auto
        , textAlign center
        , backgroundColor theme.secondaryLight
        , Css.width (px 600)
        , Css.height (Css.em 20)
        ]


inputContainer : List (Attribute msg) -> List (Html msg) -> Html msg
inputContainer =
    styled div
        [ displayFlex
        , flexWrap Css.wrap
        , alignItems stretch
        , Css.height (px 75)
        ]


urlInput : List (Attribute msg) -> List (Html msg) -> Html msg
urlInput =
    styled input
        [ border3 (px 2) solid theme.primaryDark
        , backgroundColor theme.primaryLight
        , margin (px 10)
        , padding (px 5)
        , Css.width (px 300)
        , fontFamilies [ "monospace" ]
        ]


shortBtn : List (Attribute msg) -> List (Html msg) -> Html msg
shortBtn =
    styled button
        [ border3 (px 2) solid theme.primaryDark
        , backgroundColor theme.primaryLight
        , Css.width (px 100)
        , margin (px 10)
        , fontFamilies [ "monospace" ]
        ]


view : Model -> Html Msg
view model =
    mainContainer []
        [ h1 [ css [ fontFamilies [ "monospace" ] ] ] [ text "Short URL Service" ]
        , shortApp model
        ]


shortApp : Model -> Html Msg
shortApp model =
    case model of
        Init url ->
            div []
                [ inputContainer []
                    [ urlInput [ placeholder "url to short", value url, onInput Change ] []
                    , shortBtn [ onClick (Short url) ] [ text "Short" ]
                    ]
                ]

        Loading ->
            text "Loading..."

        Failure url ->
            div []
                [ div [] [ text "I could not short given url for some reason." ]
                , inputContainer []
                    [ urlInput [ placeholder "url to short", value url, onInput Change ] []
                    , shortBtn [ onClick (Short url) ] [ text "Short" ]
                    ]
                ]

        Success shorted new ->
            let
                extendedUrl =
                    api ++ "/" ++ shorted
            in
            div []
                [ div [] [ a [ href extendedUrl ] [ text extendedUrl ] ]
                , inputContainer []
                    [ urlInput [ placeholder "url to short", value new, onInput Change ] []
                    , shortBtn [ onClick (Short new) ] [ text "Short" ]
                    ]
                ]


api : String
api =
    "http://short.beniamindudek.xyz"


apiShort : String
apiShort =
    api ++ "/api/short"


shortUrl : String -> Cmd Msg
shortUrl toShort =
    Http.post
        { body = Http.jsonBody (keyEncoder toShort)
        , expect = Http.expectJson Response keyDecoder
        , url = apiShort
        }


keyDecoder : D.Decoder String
keyDecoder =
    D.field "key" D.string


keyEncoder : String -> E.Value
keyEncoder url =
    E.object
        [ ( "value", E.string url ) ]
